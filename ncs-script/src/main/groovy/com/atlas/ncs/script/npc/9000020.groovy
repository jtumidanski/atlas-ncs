package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9000020 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int[] travelFrom = [777777777, 541000000]
   int[] travelFee = [3000, 10000]
   int[] travelMap = [800000000, 550000000]
   String[] travelPlace = ["Mushroom Shrine of Japan", "Trend Zone of Malaysia"]
   String[] travelPlaceShort = ["Mushroom Shrine", "Metropolis"]
   String[] travelPlaceCountry = ["Japan", "Malaysia"]
   String[] travelAgent = ["I", "#r#p9201135##k"]
   String[] travelDescription = ["If you desire to feel the essence of Japan, there's nothing like visiting the Shrine, a Japanese cultural melting pot. Mushroom Shrine is a mythical place that serves the incomparable Mushroom God from ancient times.",
                                 "If you desire to feel the heat of the tropics on an upbeat environment, the residents of Malaysia are eager to welcome you. Also, the metropolis itself is the heart of the local economy, that place is known to always offer something to do or to visit around."]
   String[] travelDescription2 = ["Check out the female shaman serving the Mushroom God, and I strongly recommend trying Takoyaki, Yakisoba, and other delicious food sold in the streets of Japan. Now, let's head over to #bMushroom Shrine#k, a mythical place if there ever was one.",
                                  "Once there, I strongly suggest you to schedule a visit to Kampung Village. Why? Surely you've come to know about the fantasy theme park Spooky World? No? It's simply put the greatest theme park around there, it's worth a visit! Now, let's head over to the #bTrend Zone of Malaysia#k."]
   int travelType
   int travelStatus

   def start() {
      travelStatus = getTravelingStatus(cm.getMapId())
      action((byte) 1, (byte) 0, 0)
   }

   def getTravelingStatus(int mapId) {
      for (int i = 0; i < travelMap.length; i++) {
         if (mapId == travelMap[i]) {
            return i
         }
      }

      return -1
   }

   def getTravelType(int mapId) {
      for (int i = 0; i < travelFrom.length; i++) {
         if (mapId == travelFrom[i]) {
            return i
         }
      }

      return 0
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         if (mode == 0 && status == 4) {
            status -= 2
         } else {
            cm.dispose()
            return
         }
      }

      if (travelStatus != -1) {
         if (status == 0) {
            cm.sendSimple("9000020_HOW_IS_THE_TRAVELING", cm.peekSavedLocation("WORLDTOUR"))

         } else if (status == 1) {
            if (selection == 0) {
               cm.sendNext("9000020_ALRIGHT")

            } else if (selection == 1) {
               cm.sendOk("9000020_LET_ME_KNOW")

               cm.dispose()
            }
         } else if (status == 2) {
            int map = cm.getSavedLocation("WORLDTOUR")
            if (map == -1) {
               map = 104000000
            }

            cm.warp(map)
            cm.dispose()
         }
      } else {
         if (status == 0) {
            travelType = getTravelType(cm.getMapId())
            cm.sendNext("9000020_GETTING_OUT_FOR_A_CHANGE", cm.numberWithCommas(travelFee[travelType]))

         } else if (status == 1) {
            cm.sendSimple("9000020_TRAVEL_GUIDE", travelPlace[travelType], travelAgent[travelType], travelPlaceShort[travelType], travelPlaceShort[travelType], travelPlaceCountry[travelType])

         } else if (status == 2) {
            cm.sendNext("Would you like to travel to #b" + travelPlace[travelType] + "#k? " + travelDescription[travelType])
         } else if (status == 3) {
            if (cm.getMeso() < travelFee[travelType]) {
               cm.sendNext("9000020_NOT_ENOUGH_MESOS")

               cm.dispose()
               return
            }
            cm.sendNextPrev(travelDescription2[travelType])
         } else if (status == 4) {
            cm.gainMeso(-travelFee[travelType])
            cm.saveLocation("WORLDTOUR")
            cm.warp(travelMap[travelType], 0)
            cm.dispose()
         }
      }
   }
}

NPC9000020 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9000020(cm: cm))
   }
   return (NPC9000020) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }