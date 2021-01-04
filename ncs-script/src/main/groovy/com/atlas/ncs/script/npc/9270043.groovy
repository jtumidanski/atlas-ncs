package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9270043 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int[] ids = [2000004, 2020012, 2000005, 2030007, 2022027, 2040001, 2041002, 2040805, 2040702, 2043802, 2040402, 2043702, 1302022, 1322021, 1322026, 1302026, 1442017, 1082147, 1102043, 1442016, 1402012, 1302027, 1322027, 1322025, 1312012, 1062000, 1332020, 1302028, 1372002, 1002033, 1092022, 1302021, 1102041, 1102042, 1322024, 1082148, 1002012, 1322012, 1322022, 1002020, 1302013, 1082146, 1442014, 1002096, 1302017, 1442012, 1322010, 1442011, 1442018, 1092011, 1092014, 1302003, 1432001, 1312011, 1002088, 1041020, 1322015, 1442004, 1422008, 1302056, 1432000, 1382001, 1041053, 1060014, 1050053, 1051032, 1050073, 1061036, 1002253, 1002034, 1051025, 1050067, 1051052, 1002072, 1002144, 1051054, 1050069, 1372007, 1050056, 1050074, 1002254, 1002274, 1002218, 1051055, 1382010, 1002246, 1050039, 1382007, 1372000, 1002013, 1050072, 1002036, 1002243, 1372008, 1382008, 1382011, 1092021, 1051034, 1050047, 1040019, 1041031, 1051033, 1002153, 1002252, 1051024, 1002153, 1050068, 1382003, 1382006, 1050055, 1051031, 1050025, 1002155, 1002245, 1452004, 1452023, 1060057, 1040071, 1002137, 1462009, 1452017, 1040025, 1041027, 1452005, 1452007, 1061057, 1472006, 1472019, 1060084, 1472028, 1002179, 1082074, 1332015, 1432001, 1060071, 1472007, 1472002, 1051009, 1061037, 1332016, 1332034, 1472020, 1102084, 1102086, 1102042, 1032026, 1082149]

   def start() {
      if (cm.haveItem(5451000)) {
         cm.gainItem(5451000, (short) -1)
         cm.doGachapon()
         cm.dispose()
      } else if (cm.haveItem(5220000)) {
         cm.sendYesNo("9270043_USE_YOUR_GACHAPON_TICKET")
      } else {
         cm.sendSimple("9270043_WELCOME", cm.getMapName())
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1 && cm.haveItem(5220000)) {
         cm.doGachapon()
         cm.dispose()
      } else {
         if (mode > 0) {
            status++
            if (selection == 0) {
               cm.sendNext("9270043_GACHAPON_INFO")
            } else if (selection == 1) {
               cm.sendNext("9270043_TICKET_INFO")
               cm.dispose()
            } else if (status == 2) {
               cm.sendNext("9270043_VARIETY_OF_ITEMS", cm.getMapName(), cm.getMapName())
               cm.dispose()
            }
         }
      }
   }

}

NPC9270043 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9270043(cm: cm))
   }
   return (NPC9270043) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }