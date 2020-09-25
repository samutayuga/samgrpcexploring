package cqlimpl

import (
	"context"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	cqlcrude "github.com/samutayuga/samgrpcexploring/cqlcrude/cqlcrudepb"
	"github.com/samutayuga/samgrpcexploring/sandra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	//DevKS ...
	DevKS = "dev_ota_devices"
	//IMEILatestTabl ...
	IMEILatestTabl = "ota_devices_latest_imeievent_by_iccid"
	//CID ...
	CID = "iccid"
	//CEO ...
	CEO = "event_origin"
	//CET ...
	CET = "event_type"
	//CIMSI ...
	CIMSI = "imsi"
	//CMSISDN ...
	CMSISDN = "msisdn"
	//CNEWIME ...
	CNEWIME = "new_imei"
	//CNS ...
	CNS = "notify_status"
	//COPN ...
	COPN = "operator_name"
	//COI ...
	COI = "old_imei"
	//CSEQNUM ...
	CSEQNUM = "sequence_num"
	//CSE ...
	CSE = "service_engine"
	//CTACID ...
	CTACID = "tac_id"
	//CTP ...
	CTP = "terminal_profile"
	//CTRACKST ...
	CTRACKST = "tracking_status"
	//CTRACTTM ...
	CTRACTTM = "tracking_time"
)

var (
	eventInsertCql = ""
	eventListCql   = "select * from %s"
	eventDeleteCql = "delete from %s where iccid=?"
	eventGetCql    = "select * from %s where iccid=?"
)

//CQLService ...
type CQLService struct {
}

//BuildInsertStmt ...
func BuildInsertStmt(t string, c ...string) string {
	s := fmt.Sprintf("insert into %s", t)
	ph := "values("
	for idx, col := range c {
		if idx == 0 {
			s = fmt.Sprintf("%s(%s", s, col)
			ph = fmt.Sprintf("%s%s", ph, "?")
		} else {
			if idx < len(c)-1 {
				s = fmt.Sprintf("%s,%s", s, col)
				ph = fmt.Sprintf("%s,%s", ph, "?")
			} else {
				s = fmt.Sprintf("%s,%s)", s, col)
				ph = fmt.Sprintf("%s,%s)", ph, "?")
			}

		}

	}
	s = fmt.Sprintf("%s%s", s, ph)
	return s
}

//BuildSelectStmt ...
func BuildSelectStmt(t string, w []string, c ...string) string {
	s := "select"
	//ph := "values("
	for idx, col := range c {
		if idx == 0 {
			s = fmt.Sprintf("%s %s", s, col)
		} else {
			if idx < len(c)-1 {
				s = fmt.Sprintf("%s,%s", s, col)
			} else {
				s = fmt.Sprintf("%s,%s from %s", s, col, t)
			}

		}

	}
	for i, wc := range w {
		if i == 0 {
			s = fmt.Sprintf("%s WHERE %s=?", s, wc)
		} else {
			//if i < len(w)-1 {
			s = fmt.Sprintf("%s AND %s=?", s, wc)
			//}
		}
	}
	return s
}

//SubmitTrackingEvent insert the tracking event into table
//iccid,
//event_origin,
//event_type,
//imsi,msisdn,
//new_imei,notify_status,old_imei,
//operator_name,sequence_num,service_engine,
//tac_id,terminal_profile,
//tracking_status,
//tracking_time
func (s *CQLService) SubmitTrackingEvent(ctx context.Context, req *cqlcrude.SubmitTrackingEventRequest) (*cqlcrude.SubmitTrackingEventResponse, error) {

	//request read
	evt := req.GetTrackingEvent()
	if err := sandra.Csess.Query(eventInsertCql,
		evt.GetIccid(),           //1
		evt.GetEventOrigin(),     //2
		evt.GetEventType(),       //3
		evt.GetImsi(),            //4
		evt.GetMsisdn(),          //5
		evt.GetNewImei(),         //6
		evt.GetNotifyStatus(),    //7
		evt.GetOldImei(),         //8
		evt.GetOperatorName(),    //9
		evt.GetSequenceNum(),     //10
		evt.GetServiceEngine(),   //11
		evt.GetTacId(),           //12
		evt.GetTerminalProfile(), //13
		evt.GetTrackingStatus(),  //14
		evt.GetTrackingTime()).Exec(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while inserting an event with iccid= %s,tracking time=%d,detail error: %v", evt.GetIccid(), evt.GetTrackingTime(), err))
		//log.Fatalf("Error while inserting a blog with author= %s,title=%s, %v", bRaw.GetAuthorId(), bRaw.GetTitle(), err)
	}
	return &cqlcrude.SubmitTrackingEventResponse{SubmitStatus: 0}, nil
}

//GetTrackingEvent ...
func (s *CQLService) GetTrackingEvent(ctx context.Context, req *cqlcrude.GetTrackingEventRequest) (*cqlcrude.GetTrackingEventResponse, error) {
	var errGet error
	evt := cqlcrude.TrackingEvent{}

	if errGet = sandra.Csess.Query(eventGetCql, req.GetIccid()).Consistency(gocql.One).Scan(&evt.Iccid, //1
		&evt.EventOrigin,     //2
		&evt.EventType,       //3
		&evt.Imsi,            //4
		&evt.Msisdn,          //5
		&evt.NewImei,         //6
		&evt.NotifyStatus,    //7
		&evt.OldImei,         //8
		&evt.OperatorName,    //9
		&evt.SequenceNum,     //10
		&evt.ServiceEngine,   //11
		&evt.TacId,           //12
		&evt.TerminalProfile, //13
		&evt.TrackingStatus,  //14
		&evt.TrackingTime); errGet == nil {
		return &cqlcrude.GetTrackingEventResponse{TrackingEvent: &evt}, nil

	}
	if errGet.Error() == "not found" {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Tracking event with iccid=%v is not found", req.GetIccid()))
	}
	return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while retrieving id=%v.Server Error: %v", req.GetIccid(), errGet))
}

//ListAllTrackingEvents ...
func (s *CQLService) ListAllTrackingEvents(req *cqlcrude.ListTrackingEventRequest, stream cqlcrude.CrudeService_ListAllTrackingEventsServer) error {
	iter := sandra.Csess.Query(eventListCql).Iter()
	defer iter.Close()
	evt := cqlcrude.TrackingEvent{}
	for iter.Scan(&evt.Iccid, //1
		&evt.EventOrigin,     //2
		&evt.EventType,       //3
		&evt.Imsi,            //4
		&evt.Msisdn,          //5
		&evt.NewImei,         //6
		&evt.NotifyStatus,    //7
		&evt.OldImei,         //8
		&evt.OperatorName,    //9
		&evt.SequenceNum,     //10
		&evt.ServiceEngine,   //11
		&evt.TacId,           //12
		&evt.TerminalProfile, //13
		&evt.TrackingStatus,  //14
		&evt.TrackingTime) {
		stream.Send(&cqlcrude.ListTrackingEventResponse{TrackingEvent: &evt})
	}
	return nil
}

//DeleteTrackingEvent ...
func (s *CQLService) DeleteTrackingEvent(ctx context.Context, req *cqlcrude.DeleteTrackingEventRequest) (*cqlcrude.DeleteTrackingEventResponse, error) {
	return nil, nil
}

func init() {
	//build insert statement
	//EventInsertCql : iccid ascii PRIMARY KEY,
	//event_origin ascii,
	//event_type ascii,
	//imsi ascii,
	//msisdn ascii,
	//new_imei ascii,
	//notify_status ascii,
	//old_imei ascii,
	//operator_name ascii,
	//sequence_num ascii,
	//service_engine ascii,
	//tac_id ascii,
	//terminal_profile ascii,
	//tracking_status ascii,
	//tracking_time timestamp
	//table name ota_devices_latest_imeievent_by_iccid
	eventInsertCql = BuildInsertStmt(IMEILatestTabl,
		CID,
		CEO,
		CET,
		CIMSI,
		CMSISDN,
		CNEWIME,
		CNS,
		COI,
		COPN,
		CSEQNUM,
		CSE,
		CTACID,
		CTP,
		CTRACKST,
		CTRACTTM,
	)
	log.Printf("Insert CQL %s ", eventInsertCql)
	eventGetCql = BuildSelectStmt(IMEILatestTabl, []string{"iccid"},
		CID,
		CEO,
		CET,
		CIMSI,
		CMSISDN,
		CNEWIME,
		CNS,
		COI,
		COPN,
		CSEQNUM,
		CSE,
		CTACID,
		CTP,
		CTRACKST,
		CTRACTTM)
	log.Printf("GET CQL %s ", eventGetCql)
	eventDeleteCql = fmt.Sprintf(eventDeleteCql, IMEILatestTabl)
	log.Printf("Delete CQL %s ", eventDeleteCql)
	eventListCql = BuildSelectStmt(IMEILatestTabl, []string{},
		CID,
		CEO,
		CET,
		CIMSI,
		CMSISDN,
		CNEWIME,
		CNS,
		COI,
		COPN,
		CSEQNUM,
		CSE,
		CTACID,
		CTP,
		CTRACKST,
		CTRACTTM)
	log.Printf("List CQL %s ", eventListCql)

}
