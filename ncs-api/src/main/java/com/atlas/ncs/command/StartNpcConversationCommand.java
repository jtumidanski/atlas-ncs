package com.atlas.ncs.command;

public record StartNpcConversationCommand(int worldId, int channelId, int mapId, int characterId, int npcId, int npcObjectId) {
}
