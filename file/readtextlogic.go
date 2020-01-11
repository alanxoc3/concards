package file

func greaterthan_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: fallthrough
      case OnGroup:
         resetGroup(info)
         resetCard(info)
      case OnQuestion: info.shouldCreateCard = true
   }
}

func lessthan_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break
      case OnQuestion: info.shouldCreateCard = true
   }
}

func exclude_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: fallthrough
      case OnQuestion: info.excludes = map[string]bool{}
   }
}

func question_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break
      case OnQuestion: info.shouldCreateCard = true
   }
}

func answer_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // do nothing for answers following groups (for now)
      case OnQuestion: info.answers = append(info.answers, []string{})
   }
}

func note_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: notes for group in future?
      case OnQuestion: info.notes = append(info.notes, []string{})
   }
}

func timestamp_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: meta for group in future?
      case OnQuestion: info.timestamp = []string{}
   }
}

func meta_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: meta for group in future?
      case OnQuestion: info.meta = []string{}
   }
}

// There is a forward looking approach, and there is a backwards approach.
// Group and exclude are both already doing the forward approach. I like it
// too. I'll try it. I think I did the backwards approach for the previous one.
