package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1002003 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (status == 0 && mode == 0) {
            cm.sendNext("1002003_COME_BACK_AND_TALK_BUSINESS")
            cm.dispose()
            return
         } else if (status >= 1 && mode == 0) {
            cm.sendNext("1002003_FINANCIAL_RELIEF")
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0) {
            cm.sendYesNo("1002003_DO_YOU_WANT_TO")
         } else if (status == 1) {
            cm.sendYesNo("1002003_OK_GOOD")
         } else if (status == 2) {
            def capacity = cm.getBuddyListCapacity()
            if (capacity >= 50 || cm.getMeso() < 240000) {
               cm.sendNext("1002003_ARE_YOU_SURE_YOU_HAVE_MONEY")
               cm.dispose()
            } else {
               int newCapacity = capacity + 5
               boolean status = cm.updateBuddyCapacity(newCapacity)
               if (status) {
                  cm.gainMeso(-240000)
                  cm.sendOk("1002003_SUCCESS")
               } else {
                  cm.sendOk("1002003_ISSUE_TRY_AGAIN")
               }
               cm.dispose()
            }
         }
      }
   }
}

NPC1002003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1002003(cm: cm))
   }
   return (NPC1002003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }