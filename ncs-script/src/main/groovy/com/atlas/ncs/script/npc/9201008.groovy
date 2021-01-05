package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201008 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int wid
   boolean isMarrying

   boolean cathedralWedding = false
   String weddingEventName = "WeddingChapel"
   int weddingEntryTicketCommon = 5251001
   int weddingEntryTicketPremium = 5251002
   int weddingSendTicket = 4031377
   int weddingGuestTicket = 4031406
   int weddingAltarMapId = 680000110
   boolean weddingIndoors

   static def isWeddingIndoors(mapId) {
      return mapId >= 680000100 && mapId <= 680000500
   }

   def hasSuitForWedding(int characterId) {
      int baseId = cm.characterGender(characterId) ? 1050131 : 1051150
      for (int i = 0; i < 4; i++) {
         if (cm.characterHasItem(characterId, baseId + i, true)) {
            return true
         }
      }

      return false
   }

   def getMarriageInstance(weddingId) {
      EventManager em = cm.getEventManager(weddingEventName).orElseThrow()
      for (Iterator<EventInstanceManager> iterator = em.getInstances().iterator(); iterator.hasNext();) {
         EventInstanceManager eim = iterator.next()
         if (eim.getIntProperty("weddingId") == weddingId) {
            return eim
         }
      }

      return null
   }

   def hasWeddingRing(int characterId) {
      int[] rings = [1112806, 1112803, 1112807, 1112809]
      for (int i = 0; i < rings.length; i++) {
         if (cm.characterHasItem(characterId, rings[i], true)) {
            return true
         }
      }

      return false
   }

   def start() {
      weddingIndoors = isWeddingIndoors(cm.getMapId())
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

         if (!weddingIndoors) {
            boolean hasEngagement = false
            for (int x = 4031357; x <= 4031364; x++) {
               if (cm.haveItem(x, 1)) {
                  hasEngagement = true
                  break
               }
            }

            if (status == 0) {
               String text = "Welcome to the #bChapel#k! How can I help you?"
               String[] choice = ["How do I prepare a wedding?", "I have an engagement and want to arrange the wedding", "I am the guest and I'd like to go into the wedding"]
               for (int x = 0; x < choice.length; x++) {
                  text += "\r\n#L" + x + "##b" + choice[x] + "#l"
               }

               if (cm.haveItem(5251100)) {
                  text += "\r\n#L" + x + "##bMake additional invitation cards#l"
               }

               cm.sendSimple(text)
            } else if (status == 1) {
               World world = cm.getClient().getWorldServer()
               Channel channel = cm.getClient().getChannelServer()
               switch (selection) {
                  case 0:
                     cm.sendOk("9201008_ENGAGEMENT_INFO", weddingEntryTicketCommon)
                     cm.dispose()
                     break
                  case 1:
                     if (hasEngagement) {
                        int weddingId = world.getRelationshipId(cm.getCharacterId())
                        if (weddingId > 0) {
                           if (channel.isWeddingReserved(weddingId)) {    // registration check
                              String placeTime = channel.getWeddingReservationTimeLeft(weddingId)
                              cm.sendOk("9201008_WEDDING_TIME", placeTime)

                           } else {
                              if (!cm.hasPartner()) {
                                 cm.sendOk("9201008_PARTNER_OFFLINE")
                                 cm.dispose()
                                 return
                              }

                              if (hasWeddingRing(cm.getCharacterId()) || hasWeddingRing(cm.getPartnerId())) {
                                 cm.sendOk("9201008_YOU_OR_PARTNER_HAS_MARRIAGE_RING_ALREADY")
                                 cm.dispose()
                                 return
                              }

                              if (!cm.partnerInMap()) {
                                 cm.sendOk("9201008_PARNTER_MUST_REGISTER")
                                 cm.dispose()
                                 return
                              }

                              if (!cm.canHold(weddingSendTicket, 15) || !cm.characterCanHold(cm.getPartnerId(), weddingSendTicket, 15)) {
                                 cm.sendOk("9201008_YOU_OR_PARTNER_NEED_ETC_SPACE")
                                 cm.dispose()
                                 return
                              }

                              if (!cm.getUnclaimedMarriageGifts().isEmpty() || !partner.getAbstractPlayerInteraction().getUnclaimedMarriageGifts().isEmpty()) {
                                 cm.sendOk("9201008_SOMETHING_DOES_NOT_SEEM_RIGHT")
                                 cm.dispose()
                                 return
                              }

                              boolean hasCommon = cm.haveItem(weddingEntryTicketCommon)
                              boolean hasPremium = cm.haveItem(weddingEntryTicketPremium)

                              if (hasCommon || hasPremium) {
                                 boolean weddingType = hasPremium

                                 MapleCharacter player = cm.getPlayer()
                                 int resStatus = channel.pushWeddingReservation(weddingId, cathedralWedding, weddingType, player.getId(), player.getPartnerId())
                                 if (resStatus > 0) {
                                    cm.gainItem((weddingType) ? weddingEntryTicketPremium : weddingEntryTicketCommon, (short) -1)

                                    long expirationTime = WeddingProcessor.getInstance().getRelativeWeddingTicketExpireTime(resStatus)
                                    cm.gainItem(weddingSendTicket, (short) 15, false, true, expirationTime)
                                    partner.getAbstractPlayerInteraction().gainItem(weddingSendTicket, (short) 15, false, true, expirationTime)

                                    String placeTime = channel.getWeddingReservationTimeLeft(weddingId)

                                    String wedType = weddingType ? "Premium" : "Regular"
                                    cm.sendOk("9201008_RECEIVED_TICKETS", wedType, placeTime)
                                    cm.characterSendBlueText(cm.getCharacterId(), "MARRIAGE_WEDDING_ASSISTANT", wedType, placeTime)
                                    cm.characterSendBlueText(cm.getPartnerId(), "MARRIAGE_WEDDING_ASSISTANT", wedType, placeTime)

                                    if (!hasSuitForWedding(cm.getCharacterId())) {
                                       cm.characterSendPinkText(cm.getCharacterId(), "MARRIAGE_WEDDING_ASSISTANT_GARMENT")
                                    }

                                    if (!hasSuitForWedding(cm.getPartnerId())) {
                                       cm.characterSendPinkText(cm.getPartnerId(), "MARRIAGE_WEDDING_ASSISTANT_GARMENT")
                                    }
                                 } else {
                                    cm.sendOk("9201008_PLEASE_TRY_AGAIN_LATER")
                                 }
                              } else {
                                 cm.sendOk("9201008_BEFORE_TRYING_TO_REGISTER", weddingEntryTicketCommon)
                              }
                           }
                        } else {
                           cm.sendOk("9201008_ENCOUNTERED_ERROR")
                        }

                        cm.dispose()
                     } else {
                        cm.sendOk("9201008_NEED_ENGAGEMENT_RING")
                        cm.dispose()
                     }
                     break

                  case 2:
                     if (cm.haveItem(weddingGuestTicket)) {
                        wid = channel.getOngoingWedding(cathedralWedding)
                        if (wid > 0) {
                           if (channel.isOngoingWeddingGuest(cathedralWedding, cm.getCharacterId())) {
                              EventInstanceManager eim = getMarriageInstance(wid)
                              if (eim != null) {
                                 cm.sendOk("9201008_DO_NOT_DROP_GOLD_MAPLE_LEAF")
                              } else {
                                 cm.sendOk("9201008_PLEASE_WAIT")
                                 cm.dispose()
                              }
                           } else {
                              cm.sendOk("9201008_YOU_ARE_NOT_INVITED")
                              cm.dispose()
                           }
                        } else {
                           cm.sendOk("9201008_NO_WEDDING_BOOKED")
                           cm.dispose()
                        }
                     } else {
                        cm.sendOk("9201008_NEED_TICKET", weddingGuestTicket)
                        cm.dispose()
                     }
                     break
                  default:
                     int weddingId = world.getRelationshipId(cm.getCharacterId())
                     int resStatus = channel.getWeddingReservationStatus(weddingId, cathedralWedding)
                     if (resStatus > 0) {
                        if (cm.canHold(weddingSendTicket, 3)) {
                           cm.gainItem(5251100, (short) -1)

                           long expirationTime = WeddingProcessor.getInstance().getRelativeWeddingTicketExpireTime(resStatus)
                           cm.gainItem(weddingSendTicket, (short) 3, false, true, expirationTime)
                        } else {
                           cm.sendOk("9201008_NEED_ETC_SPACE")
                        }
                     } else {
                        cm.sendOk("9201008_NOT_CURRENTLY_BOOKED")
                     }
                     cm.dispose()
               }
            } else if (status == 2) {   // registering guest
               EventInstanceManager eim = getMarriageInstance(wid)

               if (eim != null) {
                  cm.gainItem(weddingGuestTicket, (short) -1)
                  eim.registerPlayer(cm.getCharacterId())     //cm.warp(680000210, 0);
               } else {
                  cm.sendOk("9201008_MARRIAGE_EVENT_COULD_NOT_BE_FOUND")
               }

               cm.dispose()
            }
         } else {
            if (status == 0) {
               EventInstanceManager eim = cm.getEventInstance()
               if (eim == null) {
                  cm.warp(680000000, 0)
                  cm.dispose()
                  return
               }

               isMarrying = (cm.getCharacterId() == eim.getIntProperty("groomId") || cm.getCharacterId() == eim.getIntProperty("brideId"))

               if (eim.getIntProperty("weddingStage") == 0) {
                  if (!isMarrying) {
                     cm.sendOk("9201008_WELCOME", cm.getMapId())
                  } else {
                     cm.sendOk("9201008_WELCOME_SHORT", cm.getMapId())
                  }

                  cm.dispose()
               } else {
                  cm.sendYesNo("9201008_BRIDE_AND_GROOM_ARE_READY")
               }
            } else if (status == 1) {
               cm.warp(weddingAltarMapId, "sp")
               cm.dispose()
            }
         }
      }
   }
}

NPC9201008 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201008(cm: cm))
   }
   return (NPC9201008) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }