package algolog

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, algorun_id, event, level, message"

func Insert(algolog *proto.Algolog) error {
	algolog.Id = db.NewId("algologs")
	_, err := db.D.Handle.NamedExec("insert into algologs ("+Columns+") "+
		"values (:id, :algorun_id, :event, :level, :message)", algolog)
	if err != nil {
		log.Printf("algorun Create err: %v", err)
		return err
	}
	return nil
}
