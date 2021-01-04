package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2083006 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int[] quests = [3719, 3724, 3730, 3736, 3742, 3748]
   String[] array = ["Year 2021 - Average Town Entrance", "Year 2099 - Midnight Harbor Entrance", "Year 2215 - Bombed City Center Retail District", "Year 2216 - Ruined City Intersection", "Year 2230 - Dangerous Tower Lobby", "Year 2503 - Air Battleship Bow"/*, "Year 2227 - Dangerous City Intersection"*/]
   int limit

   def start() {
      if (!cm.isQuestCompleted(3718)) {
         cm.sendOk("2083006_NOT_BEEN_ACTIVATED")

         cm.dispose()
         return
      }

      for (limit = 0; limit < quests.length; limit++) {
         if (!cm.isQuestCompleted(quests[limit])) {
            break
         }
      }

      if (limit == 0) {
         cm.sendOk("2083006_PROVE_VALOR")

         cm.dispose()
         return
      }

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
            String menuSel = generateSelectionMenu(array, limit)
            cm.sendSimple(menuSel)
         } else if (status == 1) {
            int mapId = 0

            switch (selection) {
               case 0:
                  mapId = 240070100
                  break
               case 1:
                  mapId = 240070200
                  break
               case 2:
                  mapId = 240070300
                  break
               case 3:
                  mapId = 240070400
                  break
               case 4:
                  mapId = 240070500
                  break
               case 5:
                  mapId = 240070600
                  break
            /*case 6:
                mapId = 683070400;
                break;*/
            }

            if (mapId > 0) {
               cm.warp(mapId, 1)
            } else {
               cm.sendOk("2083006_COMPLETE_YOUR_MISSION")

            }
         }
      }
   }

   static def generateSelectionMenu(String[] array, int limit) {
      // nice tool for generating a string for the sendSimple functionality
      String menu = ""

      int len = Math.min(limit, array.length)
      for (def i = 0; i < len; i++) {
         menu += "#L" + i + "#" + array[i] + "#l\r\n"
      }
      return menu
   }
}

NPC2083006 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2083006(cm: cm))
   }
   return (NPC2083006) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }