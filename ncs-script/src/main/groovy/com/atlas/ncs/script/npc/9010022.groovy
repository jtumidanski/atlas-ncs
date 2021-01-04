package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9010022 {
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
         if (mode == 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0) {
            if (cm.getLevel() < 20) {
               cm.sendDimensionalMirror("9010022_NO_PLACE_TO_TRANSPORT")

               cm.dispose()
            } else {
               String selStr = ""
               if (cm.getLevel() >= 20 && cm.getLevel() <= 30) {
                  selStr += "#0# Ariant Coliseum"
               }

               if (cm.getLevel() >= 25) {
                  selStr += "#1# Mu Lung Dojo"
               }

               if (cm.getLevel() >= 30 && cm.getLevel() <= 50) {
                  selStr += "#2# Monster Carnival 1"
               }

               if (cm.getLevel() >= 51 && cm.getLevel() <= 70) {
                  selStr += "#3# Monster Carnival 2"
               }

               /*
               if (cm.getLevel() >= 40) { NOT IMPLEMENTED
                   selStr += "#5# Nett's Pyramid";
               }

               if (cm.getLevel() >= 25 && cm.getLevel() <= 30) { NOT IMPLEMENTED
                   selStr += "#6# Construction Site";
               }
               */

               cm.sendDimensionalMirror(selStr)
            }
         } else if (status == 1) {
            cm.saveLocation("MIRROR")
            switch (selection) {
               case 0:
                  cm.warp(980010000, 3)
                  break
               case 1:
                  cm.warp(925020000, 0)
                  break
               case 2:
                  cm.saveLocation("MONSTER_CARNIVAL")
                  cm.warp(980000000, 3)
                  break
               case 3:
                  cm.saveLocation("MONSTER_CARNIVAL")
                  cm.warp(980030000, 3)
                  break
               case 5:
                  cm.warp(926010000, 4)
                  break
               case 6:
                  cm.warp(910320000, 2)
                  break
            }
            cm.dispose()
         }
      }
   }
}

NPC9010022 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9010022(cm: cm))
   }
   return (NPC9010022) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }