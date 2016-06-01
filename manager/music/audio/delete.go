package audio

import (
  "MuShare/datatype/request/music"
  "MuShare/datatype"
  "MuShare/db/models"
)

func (this *Audio) DeleteAudio(body *music.Audio) datatype.Response {
  audio := models.Audio{}
  sheet := models.Sheet{}
  tx := this.DB.Begin()

  tx.Where("audio_url = ?", body.AudioUrl).First(&audio)
  if audio.ID == 0 {
    return badRequest("no such audio")
  }
  tx.Where("id + ?", audio.SheetID).First(&sheet)
  if body.UserID != sheet.UserID{
    return forbidden("audio belongs to the user")
  }

  tx.Delete(&audio)
  tx.Commit()

  return ok("delete ok", audio)
}