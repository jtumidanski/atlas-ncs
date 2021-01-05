package com.atlas.ncs.model;

import java.util.List;

public record Party(int leaderId, List<PartyCharacter> members) {
}
