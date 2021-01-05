package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201004 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int[] rings = [1112806, 1112803, 1112807, 1112809]
   int divorceFee = 500000
   Ring ringObj

   def getWeddingRingItemId(int characterId) {
      for (int i = 0; i < rings.length; i++) {
         if (cm.characterHasItem(characterId, rings[i], false)) {
            return rings[i]
         }
      }

      return null
   }

   def hasEquippedWeddingRing(int characterId) {
      for (int i = 0; i < rings.length; i++) {
         if (cm.characterHasItemEquipped(characterId, rings[i])) {
            return true
         }
      }

      return false
   }

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
            String[] questionStr = ["How can I engage someone?", "How can I marry?", "How can I divorce?"]

            if (!(!cm.isMarried() && getWeddingRingItemId(cm.getCharacterId()))) {
               questionStr << "I want a divorce..."
            } else {
               questionStr << "I wanna remove my old wedding ring..."
            }

            cm.sendSimple("Hello, welcome to #bAmoria#k, a beautiful land where maplers can find love and, if inspired enough, even marry. Do you have any questions about Amoria? Talk it to me.#b\r\n\r\n" + generateSelectionMenu(questionStr))
         } else if (status == 1) {
            switch (selection) {
               case 0:
                  cm.sendOk("9201004_PROCESS_IS_STRAIGHT_FORWARD")
                  cm.dispose()
                  break
               case 1:
                  cm.sendOk("9201004_MUST_ALREADY_BE_ENGAGED")
                  cm.dispose()
                  break
               case 2:
                  cm.sendOk("9201004_DIVORCE_IS_POSSIBLE", divorceFee)
                  cm.dispose()
                  break
               case 3:
                  ringObj = cm.getPlayer().getMarriageRing()
                  if (ringObj == null) {
                     Object itemId = getWeddingRingItemId(cm.getCharacterId())

                     if (itemId != null) {
                        cm.sendOk("9201004_DIVOCE_SUCCESS")
                        cm.gainItem((int) itemId, (short) -1)
                     } else if (hasEquippedWeddingRing(cm.getCharacterId())) {
                        cm.sendOk("9201004_TAKE_RING_OFF")
                     } else {
                        cm.sendOk("9201004_YOU_ARE_NOT_MARRIED")
                     }

                     cm.dispose()
                     return
                  }

                  cm.sendYesNo("9201004_DIVORCE_CONFIRMATION")
                  break
            }
         } else if (status == 2) {
            if (cm.getMeso() < divorceFee) {
               cm.sendOk("9201004_NEED_DIVORCE_FEE", divorceFee)
               cm.dispose()
               return
            } else if (ringObj.equipped()) {
               cm.sendOk("9201004_TAKE_OFF_YOUR_RING")
               cm.dispose()
               return
            }

            cm.gainMeso(-divorceFee)
            RingActionHandler.breakMarriageRing(cm.getPlayer(), ringObj.itemId())
            cm.gainItem(ringObj.itemId(), (short) -1)
            cm.sendOk("9201004_DIVORCED_YOUR_PARTNER")
            cm.dispose()
         }
      }
   }

   static def generateSelectionMenu(String[] array) {
      String menu = ""
      for (int i = 0; i < array.length; i++) {
         menu += "#L" + i + "#" + array[i] + "#l\r\n"
      }
      return menu
   }
}

NPC9201004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201004(cm: cm))
   }
   return (NPC9201004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }