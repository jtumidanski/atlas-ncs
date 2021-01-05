package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC2112007 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      EventInstanceManager eim = cm.getEventInstance()
      String book = "stg1_b" + (cm.getNpcObjectId() % 26)

      int res = eim.getIntProperty(book)
      if (res > -1) {
         eim.setIntProperty(book, -1)

         if (res == 0) {  // mesos
            int mesoGain = 500 * cm.getMesoRate()
            cm.sendNext("2112007_EARNED_MESOS", mesoGain)

            cm.gainMeso(mesoGain)
         } else if (res == 1) {  // exp
            int expGain = 500 * cm.getExpRate()
            cm.sendNext("2112007_EARNED_EXP", expGain)
            cm.gainExp(expGain)
         } else if (res == 2) {  // letter
            int letter = 4001131
            if (!cm.canHold(letter)) {
               cm.sendOk("2112007_NO_INVENTORY_SPACE")
               cm.dispose()
               return
            }

            cm.gainItem(letter, (short) 1)
            cm.sendNext("2112007_FOUND_A_LETTER")
         } else if (res == 3) {  // pass
            cm.sendNext("2112007_FOUND_TRIGGER")

            eim.showClearEffect()
            eim.giveEventPlayersStageReward(1)
            eim.setIntProperty("statusStg1", 1)

            cm.getMap().getReactorByName("d00").hitReactor(cm.getClient())
         }
      } else {
         cm.sendNext("2112007_NOTHING_HERE")
      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC2112007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2112007(cm: cm))
   }
   return (NPC2112007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }