package com.atlas.ncs.command;

public record ContinueNpcConversationCommand(int characterId, byte mode, byte type, int selection) {
}
