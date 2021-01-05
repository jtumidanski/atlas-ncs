package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201107 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getMapId() == 610030500) {
         cm.sendOk("9201107_UNBELIEVABLE")
         cm.dispose()
      } else if (cm.getMapId() == 610030000) {
         cm.sendOk("9201107_LEGENDARY_FAMILY_OF_HEROES")
         cm.dispose()
      } else if (cm.getMapId() == 610030510) {
         if (cm.getMapMonsterCount() == 0) {
            EventInstanceManager eim = cm.getEventInstance()
            int stgStatus = eim.getIntProperty("glpq5_room")
            int jobNiche = cm.getJobNiche()

            if ((stgStatus >> jobNiche) % 2 == 0) {
               if (cm.canHold(4001259, 1)) {
                  cm.gainItem(4001259, (short) 1)
                  cm.sendOk("9201107_GOOD_JOB")

                  stgStatus += (1 << jobNiche)
                  eim.setIntProperty("glpq5_room", stgStatus)
               } else {
                  cm.sendOk("9201107_MAKE_ETC_ROOM")
               }
            } else {
               cm.sendOk("9201107_WEAPON_INSIDE")
            }
         } else {
            cm.sendOk("9201107_ELIMINATE_ALL")
         }
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC9201107 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201107(cm: cm))
   }
   return (NPC9201107) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }