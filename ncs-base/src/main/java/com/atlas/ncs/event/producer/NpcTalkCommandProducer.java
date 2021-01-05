package com.atlas.ncs.event.producer;

import com.atlas.csrv.command.NpcTalkCommand;
import com.atlas.csrv.command.NpcTalkSpeaker;
import com.atlas.csrv.command.NpcTalkType;
import com.atlas.csrv.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class NpcTalkCommandProducer {
   private NpcTalkCommandProducer() {
   }

   public static void sendSimple(int characterId, int npcId, String message) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.SIMPLE, NpcTalkSpeaker.NPC_LEFT));
   }

   public static void sendYesNo(int characterId, int npcId, String message) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.YES_NO, NpcTalkSpeaker.NPC_LEFT));
   }

   public static void sendAcceptDecline(int characterId, int npcId, String message) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.ACCEPT_DECLINE, NpcTalkSpeaker.NPC_LEFT));
   }

   public static void sendOk(int characterId, int npcId, String message) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.OK, NpcTalkSpeaker.NPC_LEFT));
   }

   public static void sendPrevious(int characterId, int npcId, String message) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.PREVIOUS, NpcTalkSpeaker.NPC_LEFT));
   }

   public static void sendNextPrevious(int characterId, int npcId, String message) {
      sendNextPrevious(characterId, npcId, message, NpcTalkSpeaker.NPC_LEFT.getValue());
   }

   public static void sendNextPrevious(int characterId, int npcId, String message, byte speaker) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.NEXT_PREVIOUS, NpcTalkSpeaker.fromValue(speaker)));
   }

   public static void sendNext(int characterId, int npcId, String message) {
      sendNextPrevious(characterId, npcId, message, NpcTalkSpeaker.CHARACTER_LEFT.getValue());
   }

   public static void sendNext(int characterId, int npcId, String message, byte speaker) {
      EventProducerRegistry.getInstance().send(NpcTalkCommand.class, CommandConstants.TOPIC_NPC_TALK_COMMAND, characterId,
            new NpcTalkCommand(characterId, npcId, message, NpcTalkType.NEXT, NpcTalkSpeaker.fromValue(speaker)));
   }
}
