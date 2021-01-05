package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201010 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1 || mode == 0) {
         cm.dispose()
         return
      } else if (mode == 1) {
         status++
      } else {
         status--
      }

      EventInstanceManager eim = cm.getEventInstance()
      if (eim == null) {
         cm.warp(680000000, 0)
         cm.dispose()
         return
      }

      boolean isMarrying = (cm.getCharacterId() == eim.getIntProperty("groomId") || cm.getCharacterId() == eim.getIntProperty("brideId"))

      switch (status) {
         case 0:
            if (cm.getMapId() == 680000300) {
               cm.sendYesNo("9201010_ARE_YOU_SURE")

            } else {
               boolean hasEngagement = false
               for (int x = 4031357; x <= 4031364; x++) {
                  if (cm.haveItem(x, 1)) {
                     hasEngagement = true
                     break
                  }
               }

               if (cm.haveItem(4000313) && isMarrying) {
                  if (eim.getIntProperty("weddingStage") == 3) {
                     cm.sendOk("9201010_TOTALLY_ROCKED")

                     cm.dispose()
                  } else if (hasEngagement) {
                     if (!cm.createMarriageWishList()) {
                        cm.sendOk("9201010_ALREADY_SENT_WISH_LIST")

                     }
                     cm.dispose()
                  } else {
                     cm.sendOk("9201010_OH_HEY")

                  }
               } else {
                  if (eim.getIntProperty("weddingStage") == 3) {
                     if (!isMarrying) {
                        cm.sendYesNo("9201010_THEY_WILL_START_SOON")
                     } else {
                        cm.sendOk("9201010_TOTALLY_ROCKED")
                        cm.dispose()
                     }
                  } else {
                     cm.sendYesNo("9201010_SKIPPING_BONUS")
                  }
               }
            }

            break
         case 1:
            cm.warp(680000000, 0)
            cm.dispose()
            break
      }
   }
}

NPC9201010 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201010(cm: cm))
   }
   return (NPC9201010) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }