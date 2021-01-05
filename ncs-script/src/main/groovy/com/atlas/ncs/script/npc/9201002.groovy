package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201002 {
   NPCConversationManager cm
   int status = -1
   int sel = -1
   int state
   EventInstanceManager eim
   String weddingEventName = "WeddingCathedral"
   boolean cathedralWedding = true
   boolean weddingIndoors
   int weddingBlessingExp = cm.getConfiguration().weddingBlessingExperience()

   static def isWeddingIndoors(int mapId) {
      return mapId >= 680000100 && mapId <= 680000500
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

   def detectPlayerItemId(int characterId) {
      for (int x = 4031357; x <= 4031364; x++) {
         if (cm.characterHasItem(characterId, x)) {
            return x
         }
      }

      return -1
   }

   static def getRingId(boxItemId) {
      return boxItemId == 4031357 ? 1112803 : (boxItemId == 4031359 ? 1112806 : (boxItemId == 4031361 ? 1112807 : (boxItemId == 4031363 ? 1112809 : -1)))
   }

   def isSuitedForWedding(int characterId, boolean equipped) {
      int baseId = cm.characterGender(characterId) == 0 ? 1050131 : 1051150

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

   def getWeddingPreparationStatus(int characterId, int partnerId) {
      if (!cm.characterHasItem(characterId, 4000313)) {
         return -3
      }
      if (!cm.characterHasItem(partnerId, 4000313)) {
         return 3
      }

      if (!isSuitedForWedding(characterId, true)) {
         return -4
      }
      if (!isSuitedForWedding(partnerId, true)) {
         return 4
      }

      boolean hasEngagement = false
      for (int x = 4031357; x <= 4031364; x++) {
         if (cm.characterHasItem(characterId, x)) {
            hasEngagement = true
            break
         }
      }
      if (!hasEngagement) {
         return -1
      }

      hasEngagement = false
      for (int x = 4031357; x <= 4031364; x++) {
         if (cm.characterHasItem(partnerId, x)) {
            hasEngagement = true
            break
         }
      }
      if (!hasEngagement) {
         return -2
      }

      if (!cm.characterCanHold(characterId, 1112803)) {
         return 1
      }
      if (!cm.characterCanHold(partnerId, 1112803)) {
         return 2
      }

      return 0
   }

   def giveCoupleBlessings(EventInstanceManager eim, int characterId, int partnerId) {
      int blessCount = eim.gridSize()
      cm.characterGainExp(characterId, blessCount * weddingBlessingExp)
      cm.characterGainExp(partnerId, blessCount * weddingBlessingExp)
   }

   def start() {
      weddingIndoors = isWeddingIndoors(cm.getMapId())
      if (weddingIndoors) {
         eim = cm.getEventInstance()
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

         if (!weddingIndoors) {
            if (status == 0) {
               boolean hasEngagement = false
               for (int x = 4031357; x <= 4031364; x++) {
                  if (cm.haveItem(x, 1)) {
                     hasEngagement = true
                     break
                  }
               }

               if (hasEngagement) {
                  String text = "Hi there. How can I help you?"
                  String[] choice = ["We're ready to get married."]
                  for (int x = 0; x < choice.length; x++) {
                     text += "\r\n#L" + x + "##b" + choice[x] + "#l"
                  }
                  cm.sendSimple(text)
               } else {
                  cm.sendOk("9201002_TODAY_TWO")

                  cm.dispose()
               }
            } else if (status == 1) {
               int wid = cm.getClient().getWorldServer().getRelationshipId(cm.getCharacterId())
               Channel channel = cm.getClient().getChannelServer()

               if (channel.isWeddingReserved(wid)) {
                  if (wid == channel.getOngoingWedding(cathedralWedding)) {
                     if (!(!cm.hasPartner() || !cm.partnerInMap())) {
                        if (!cm.canHold(4000313)) {
                           cm.sendOk("9201002_ETC_SPACE_NEEDED")
                           cm.dispose()
                           return
                        } else if (!cm.characterCanHold(cm.getPartnerId(), 4000313)) {
                           cm.sendOk("9201002_PARTNER_NEEDS_ETC_SPACE")
                           cm.dispose()
                           return
                        } else if (!isSuitedForWedding(cm.getCharacterId(), false)) {
                           cm.sendOk("9201002_PURCHASE_GARMENT")
                           cm.dispose()
                           return
                        } else if (!isSuitedForWedding(cm.getPartnerId(), false)) {
                           cm.sendOk("9201002_PARTNER_GARMENT_PURCHASE")
                           cm.dispose()
                           return
                        }

                        cm.sendOk("9201002_VERY_WELL_LONG")
                     } else {
                        cm.sendOk("9201002_MISSING_PARTNER")
                        cm.dispose()
                     }
                  } else {
                     String placeTime = channel.getWeddingReservationTimeLeft(wid)

                     cm.sendOk("9201002_HAVE_PATIENCE", placeTime)
                     cm.dispose()
                  }
               } else {
                  cm.sendOk("9201002_NO_RESERVATIONS_MADE")
                  cm.dispose()
               }
            } else if (status == 2) {
               Channel channel = cm.getClient().getChannelServer()
               boolean weddingType = channel.getOngoingWeddingType(cathedralWedding)

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
                              eim.setIntProperty("isPremium", weddingType ? 1 : 0)
                              eim.registerPlayer(cm.getPartnerId())
                           } else {
                              cm.sendOk("9201002_UNEXPECTED_ERROR")
                           }

                           cm.dispose()
                        } else {
                           cm.sendOk("9201002_UNEXPECTED_ERROR_BEFORE_PREPARATIONS")
                           cm.dispose()
                        }
                     } else {
                        cm.sendOk("9201002_UNEXPECTED_ERROR_BEFORE_PREPARATIONS")
                        cm.dispose()
                     }
                  } else {    // partner already decided to start
                     cm.dispose()
                  }
               } else {
                  cm.sendOk("9201002_MISSING_PARTNER")
                  cm.dispose()
               }
            }
         } else {
            if (status == 0) {
               if (eim == null) {
                  cm.warp(680000000, 0)
                  cm.dispose()
                  return
               }

               int playerId = cm.getCharacterId()
               if (playerId == eim.getIntProperty("groomId") || playerId == eim.getIntProperty("brideId")) {
                  int weddingStage = eim.getIntProperty("weddingStage")

                  if (weddingStage == 2) {
                     cm.sendYesNo("9201002_VERY_WELL")
                     state = 1
                  } else if (weddingStage == 1) {
                     cm.sendOk("9201002_TIME_OF_HAPPINESS")
                     cm.dispose()
                  } else {
                     cm.sendOk("9201002_CONGRATULATIONS")
                     cm.dispose()
                  }
               } else {
                  int weddingStage = eim.getIntProperty("weddingStage")
                  if (weddingStage == 1) {
                     if (eim.gridCheck(cm.getCharacterId()) != -1) {
                        cm.sendOk("9201002_GIVE_YOUR_BLESSINGS")
                        cm.dispose()
                     } else {
                        if (eim.getIntProperty("guestBlessings") == 1) {
                           cm.sendYesNo("9201002_DO_YOU_WANT_TO_BLESS")
                           state = 0
                        } else {
                           cm.sendOk("9201002_TODAY_WE_ARE_GATHERED")
                           cm.dispose()
                        }
                     }
                  } else if (weddingStage == 3) {
                     cm.sendOk("9201002_NOW_MARRIED")
                     cm.dispose()
                  } else {
                     cm.sendOk("9201002_BLESSING_TIME_ENDED")
                     cm.dispose()
                  }
               }
            } else if (status == 1) {
               if (state == 0) {    // give player blessings
                  eim.gridInsert(cm.getCharacterId(), 1)

                  if (YamlConfig.config.server.WEDDING_BLESSER_SHOWFX) {
                     MapleCharacter target = cm.getPlayer()
                     PacketCreator.announce(target, new ShowSpecialEffect(9))
                     MasterBroadcaster.getInstance().sendToAllInMap(target.getMap(), new ShowForeignEffect(target.getId(), 9), false, target)
                  } else {
                     MapleCharacter target = eim.getPlayerById(eim.getIntProperty("groomId"))
                     PacketCreator.announce(target, new ShowSpecialEffect(9))
                     MasterBroadcaster.getInstance().sendToAllInMap(target.getMap(), new ShowForeignEffect(target.getId(), 9), false, target)

                     target = eim.getPlayerById(eim.getIntProperty("brideId"))
                     PacketCreator.announce(target, new ShowSpecialEffect(9))
                     MasterBroadcaster.getInstance().sendToAllInMap(target.getMap(), new ShowForeignEffect(target.getId(), 9), false, target)
                  }

                  cm.sendOk("9201002_WHAT_A_NOBLE_ACT")
                  cm.dispose()
               } else {            // couple wants to complete the wedding
                  int weddingStage = eim.getIntProperty("weddingStage")

                  if (weddingStage == 2) {
                     int pid = cm.getPartnerId()
                     if (pid <= 0) {
                        cm.sendOk("9201002_NO_LONGER_ENGAGED")
                        cm.dispose()
                        return
                     }

                     if (cm.hasPartner()) {
                        state = getWeddingPreparationStatus(cm.getCharacterId(), cm.getPartnerId())

                        switch (state) {
                           case 0:
                              pid = eim.getIntProperty("confirmedVows")
                              if (pid != -1) {
                                 if (pid == cm.getCharacterId()) {
                                    cm.sendOk("9201002_ALREADY_CONFIRMED")
                                 } else {
                                    eim.setIntProperty("weddingStage", 3)
                                    int playerItemId = detectPlayerItemId(cm.getCharacterId())
                                    int partnerItemId = (playerItemId % 2 == 1) ? playerItemId + 1 : playerItemId - 1
                                    int marriageRingId = getRingId((playerItemId % 2 == 1) ? playerItemId : partnerItemId)

                                    cm.gainItem(playerItemId, (short) -1)
                                    cm.characterGainItem(cm.getPartnerId(), partnerItemId, (short) -1)

                                    RingActionHandler.giveMarriageRings(player, partner, marriageRingId)
                                    cm.characterSetMarriageItemId(cm.getCharacterId(), marriageRingId)
                                    cm.characterSetMarriageItemId(cm.getPartnerId(), marriageRingId)

                                    //var marriageId = eim.getIntProperty("weddingId");
                                    //player.announce(Wedding.OnMarriageResult(marriageId, player, true));
                                    //partner.announce(Wedding.OnMarriageResult(marriageId, player, true));

                                    giveCoupleBlessings(eim, cm.getCharacterId(), cm.getPartnerId())
                                    cm.sendBlueTextToMap("MARRIAGE_WEDDING_SUCCESS")
                                    eim.schedule("showMarriedMsg", 2 * 1000)
                                 }
                              } else {
                                 eim.setIntProperty("confirmedVows", cm.getCharacterId())
                                 cm.sendBlueTextToMap("MARRIAGE_WEDDING_ONE_LAST_STEP", cm.getCharacterName())
                              }
                              break
                           case -1:
                              cm.sendOk("9201002_MISSING_RING_BOX")
                              break
                           case -2:
                              cm.sendOk("9201002_PARTNER_MISSING_RING_BOX")
                              break
                           case -3:
                              cm.sendOk("9201002_CANNOT_MARRY_YOU_WITHOUT_THAT_ITEM")
                              break
                           case -4:
                              cm.sendOk("9201002_GARMENTS_ARE_ESSENTIAL")
                              break
                           case 1:
                              cm.sendOk("9201002_MAKE_EQUIP_SPACE")
                              break
                           case 2:
                              cm.sendOk("9201002_PARTNER_MAKE_EQUIP_SPACE")
                              break
                           case 3:
                              cm.sendOk("9201002_PARTNER_CANNOT_MARRY_YOU_WITHOUT_THAT_ITEM")
                              break
                           case 4:
                              cm.sendOk("9201002_PARTNER_NOT_PROPERLY_DRESSED")
                              break
                        }

                        cm.dispose()
                     } else {
                        cm.sendOk("9201002_PARTNER_IS_NOT_HERE")
                        cm.dispose()
                     }
                  } else {
                     cm.sendOk("9201002_NOW_HUSBAND_AND_WIFE")
                     cm.dispose()
                  }
               }
            }
         }
      }
   }
}

NPC9201002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201002(cm: cm))
   }
   return (NPC9201002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }