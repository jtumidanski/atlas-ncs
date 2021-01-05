package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201012 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int state
   EventInstanceManager eim
   String weddingEventName = "WeddingChapel"
   boolean cathedralWedding = false

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def isSuitedForWedding(int characterId, equipped) {
      int baseId = (cm.characterGender(characterId) == 0) ? 1050131 : 1051150

      if (equipped) {
         for (int i = 0; i < 4; i++) {
            if (cm.characterHasItemEquipped(characterId, baseId + i)) {
               return true
            }
         }
      } else {
         for (int i = 0; i < 4; i++) {
            if (cm.characterHasItem(characterId, baseId + i, true)) {
               return true
            }
         }
      }

      return false
   }

   def getMarriageInstance(int characterId) {
      EventManager em = cm.getEventManager(weddingEventName).orElseThrow()
      for (Iterator<EventInstanceManager> iterator = em.getInstances().iterator(); iterator.hasNext();) {
         EventInstanceManager eim = iterator.next()
         if (eim.isEventLeader(characterId)) {
            return eim
         }
      }

      return null
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
            boolean hasEngagement = false
            for (int x = 4031357; x <= 4031364; x++) {
               if (cm.haveItem(x, 1)) {
                  hasEngagement = true
                  break
               }
            }

            if (hasEngagement) {
               String text = "Hi there. How about skyrocket the day with your fiancee baby~?"
               String[] choice = ["We're ready to get married."]
               for (int x = 0; x < choice.length; x++) {
                  text += "\r\n#L" + x + "##b" + choice[x] + "#l"
               }
               cm.sendSimple(text)
            } else {
               cm.sendOk("9201012_HELLO")

               cm.dispose()
            }
         } else if (status == 1) {
            int wid = cm.getClient().getWorldServer().getRelationshipId(cm.getCharacterId())
            Channel channel = cm.getClient().getChannelServer()

            if (channel.isWeddingReserved(wid)) {
               if (wid == channel.getOngoingWedding(cathedralWedding)) {
                  if (!(!cm.hasPartner() || !cm.partnerInMap())) {
                     int partnerId = cm.getPartnerId()
                     if (!cm.canHold(4000313)) {
                        cm.sendOk("9201012_NEED_ETC_SPACE")
                        cm.dispose()
                        return
                     } else if (!cm.characterCanHold(partnerId, 4000313)) {
                        cm.sendOk("9201012_PARTNER_NEED_ETC_SPACE")
                        cm.dispose()
                        return
                     } else if (!isSuitedForWedding(cm.getCharacterId(), false)) {
                        cm.sendOk("9201012_FASHIONABLE_CLOTHES")
                        cm.dispose()
                        return
                     } else if (!isSuitedForWedding(partnerId, false)) {
                        cm.sendOk("9201012_PARTNER_FASHIONABLE_CLOTHES")
                        cm.dispose()
                        return
                     }

                     cm.sendOk("9201012_ALRIGHT")
                  } else {
                     cm.sendOk("9201012_PARTNER_ELSEWHERE")
                     cm.dispose()
                  }
               } else {
                  String placeTime = channel.getWeddingReservationTimeLeft(wid)

                  cm.sendOk("9201012_WEDDING_TIME", placeTime)
                  cm.dispose()
               }
            } else {
               cm.sendOk("9201012_NO_RESERVATIONS")
               cm.dispose()
            }
         } else if (status == 2) {
            Channel channel = cm.getClient().getChannelServer()
            boolean weddingIsOnGoing = channel.getOngoingWeddingType(cathedralWedding)

            if (!(!cm.hasPartner() || !cm.partnerInMap())) {
               if (channel.acceptOngoingWedding(cathedralWedding)) {
                  int wid = cm.getClient().getWorldServer().getRelationshipId(cm.getCharacterId())
                  if (wid > 0) {
                     EventManager em = cm.getEventManager(weddingEventName).orElseThrow()
                     if (em.startInstance(cm.getCharacterId())) {
                        eim = getMarriageInstance(cm.getCharacterId())
                        if (eim != null) {
                           eim.setIntProperty("weddingId", wid)
                           eim.setIntProperty("groomId", cm.getCharacterId())
                           eim.setIntProperty("brideId", cm.getPartnerId())
                           eim.setIntProperty("isPremium", weddingIsOnGoing ? 1 : 0)
                           eim.registerPlayer(cm.getPartnerId())
                        } else {
                           cm.sendOk("9201012_UNEXPECTED_ERROR")
                        }

                        cm.dispose()
                     } else {
                        cm.sendOk("9201012_UNEXPECTED_ERROR_DURING_PREPARATIONS")
                        cm.dispose()
                     }
                  } else {
                     cm.sendOk("9201012_UNEXPECTED_ERROR_DURING_PREPARATIONS")
                     cm.dispose()
                  }
               } else {    // partner already decided to start
                  cm.dispose()
               }
            } else {
               cm.sendOk("9201012_PARTNER_ELSEWHERE")
               cm.dispose()
            }
         }
      }
   }
}

NPC9201012 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201012(cm: cm))
   }
   return (NPC9201012) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }