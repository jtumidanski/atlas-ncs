package com.atlas.ncs.event.producer;

import com.atlas.csrv.command.ServerNoticeCommand;
import com.atlas.csrv.command.ServerNoticeType;
import com.atlas.csrv.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class ServerNoticeProducer {
   private ServerNoticeProducer() {
   }

   public static void sendPinkText(int characterId, String message) {
      EventProducerRegistry.getInstance().send(ServerNoticeCommand.class, CommandConstants.TOPIC_SERVER_NOTICE_COMMAND, characterId,
            new ServerNoticeCommand(ServerNoticeType.PINK_TEXT, characterId, message));
   }

   public static void sendLightBlueText(int characterId, String message) {
      EventProducerRegistry.getInstance().send(ServerNoticeCommand.class, CommandConstants.TOPIC_SERVER_NOTICE_COMMAND, characterId,
            new ServerNoticeCommand(ServerNoticeType.LIGHT_BLUE, characterId, message));
   }

   public static void sendNotice(int characterId, String message) {
      EventProducerRegistry.getInstance().send(ServerNoticeCommand.class, CommandConstants.TOPIC_SERVER_NOTICE_COMMAND, characterId,
            new ServerNoticeCommand(ServerNoticeType.NOTICE, characterId, message));
   }
}
