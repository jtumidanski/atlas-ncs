package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2050017 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (cm.isQuestStarted(3421)) {
               int meteoriteId = cm.getNpc() - 2050014

               int progress = cm.getQuestProgressInt(3421, 1)
               if ((progress >> meteoriteId) % 2 == 0 || (progress == 63 && !cm.haveItem(4031117, 6))) {
                  if (cm.canHold(4031117, 1)) {
                     progress |= (1 << meteoriteId)

                     cm.gainItem(4031117, (short) 1)
                     cm.setQuestProgress(3421, 1, progress)
                  } else {
                     MessageBroadcaster.getInstance().sendServerNotice(cm.getPlayer(), ServerNoticeType.POP_UP, I18nMessage.from("HAVE_A_ETC_SLOT_AVAILABLE"))
                  }
               }
            }

            cm.dispose()
         }
      }
   }
}

NPC2050017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2050017(cm: cm))
   }
   return (NPC2050017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }