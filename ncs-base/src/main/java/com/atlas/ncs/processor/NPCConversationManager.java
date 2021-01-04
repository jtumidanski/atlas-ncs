package com.atlas.ncs.processor;

import com.atlas.ncs.NPCScriptRegistry;
import com.atlas.ncs.event.producer.CharacterEnableActionsProducer;
import com.atlas.ncs.model.Party;
import com.atlas.ncs.model.Pet;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;

public class NPCConversationManager {
   private final int characterId;

   private final int npcId;

   private final int npcOid;

   private final String scriptName;

   private final boolean itemScript;

   public NPCConversationManager(int characterId, int npc) {
      this(characterId, npc, -1, "", false);
   }

   public NPCConversationManager(int characterId, int npc, String scriptName) {
      this(characterId, npc, -1, scriptName, false);
   }

   public NPCConversationManager(int characterId, int npcId, int oid, String scriptName, boolean itemScript) {
      this.characterId = characterId;
      this.npcId = npcId;
      this.npcOid = oid;
      this.scriptName = scriptName;
      this.itemScript = itemScript;
   }

   public int getCharacterId() {
      return characterId;
   }

   public int getNpcId() {
      return npcId;
   }

   public void dispose() {
      NPCScriptRegistry.getInstance().dispose(characterId);
      CharacterEnableActionsProducer.enableActions(characterId);
   }

   public String evaluateToken(String token, Object... replacements) {
      return "";
   }

   public void sendSimple(String text, Object... replacements) {
   }

   public void sendNext(String token, Object... replacements) {
   }

   public void sendNextSpeaker(String token, byte speaker) {
   }

   public void sendNextPrev(String token, Object... replacements) {
   }

   public void sendNextPrevSpeaker(String token, byte speaker) {
   }

   public void sendYesNo(String token, Object... replacements) {
   }

   public void warp(int mapId, int portalId) {
   }

   public void warp(int mapId, String portalName) {
   }

   public void warp(int mapId) {
   }

   public void sendPrev(String token, Object... replacements) {
   }

   public void sendOk(String token, Object... replacements) {
   }

   public void lockUI() {
   }

   public boolean haveItem(int itemId) {
      return false;
   }

   public boolean haveItem(int itemId, int quantity) {
      return false;
   }

   public void gainItem(int itemId) {
      gainItem(itemId, (short) 1);
   }

   public void gainItem(int itemId, short quantity) {
   }

   public void gainItem(int itemId, boolean something) {
   }

   public void gainMeso(int meso) {
   }

   public int getMeso() {
      return 0;
   }

   public int getLevel() {
      return 0;
   }

   public int getJobId() {
      return 0;
   }

   public void saveLocation(String location) {
   }

   public int getBuddyListCapacity() {
      return 0;
   }

   public boolean hasItem(int itemId, int quantity) {
      return false;
   }

   public boolean hasItem(int itemId) {
      return false;
   }

   public boolean isQuestCompleted(int questId) {
      return false;
   }

   public boolean canHold(int itemId) {
      return false;
   }

   public boolean canHold(int itemId, int quantity) {
      return false;
   }

   public int petCount() {
      return 0;
   }

   public void gainCloseness(int amount) {
   }

   public boolean isQuestStarted(int questId) {
      return false;
   }

   public void startQuest(int questId) {
   }

   public boolean canGetFirstJob(int jobType) {
      return false;
   }

   public void setHair(int hair) {
   }

   public int getHair() {
      return 0;
   }

   public int getGender() {
      return 0;
   }

   public void sendStyle(String token, int[] ints) {
   }

   public void setSkin(int skin) {
   }

   public int getMapId() {
      return 0;
   }

   public boolean isQuestNotStarted(int questId) {
      return false;
   }

   public void forceCompleteQuest(int questId) {
   }

   public void changeJobById(int jobId) {
   }

   public void completeQuest(int questId) {
   }

   public void sendAcceptDecline(String token, Object... replacements) {
   }

   public void removeAll(int itemId) {
   }

   public int getSavedLocation(String location) {
      return 0;
   }

   public void earnTitle(String title) {
   }

   public void guideHint(int hint) {
   }

   public boolean haveItemWithId(int itemId, Boolean something) {
      return false;
   }

   public void showIntro(String s) {
   }

   public void setFace(int face) {
   }

   public void sendGetText(String token, Object... replacements) {
   }

   public void warpParty(int mapId) {
   }

   public void warpParty(int mapId, int portalId) {
   }

   public void cancelCPQLobby() {
   }

   public void gainFame(int amount) {
   }

   public void showEffect(String path) {
   }

   public boolean updateBuddyCapacity(int amount) {
      return false;
   }

   public void sendStorage(int itemId) {
   }

   public boolean gotPartyQuestItem(String s) {
      return false;
   }

   public void removePartyQuestItem(String s) {
   }

   public void setPartyQuestItemObtained(String s) {
   }

   public void killAllMonstersNotFriendly() {
   }

   public int getFace() {
      return 0;
   }

   public int peekSavedLocation(String location) {
      return 0;
   }

   public int getRemainingSp() {
      return 0;
   }

   public boolean isGM() {
      return false;
   }

   public int getFestivalPoints() {
      return 0;
   }

   public int getMapMonsterCount() {
      return 0;
   }

   public int getJobNiche() {
      return 0;
   }

   public double getExpRate() {
      return 0.0;
   }

   public double getMesoRate() {
      return 0;
   }

   public void sendGetNumber(String token, Object... replacements) {
      sendGetNumber(token, 1, 1, 100, replacements);
   }

   public void sendGetNumber(String token, int def, int min, int max, Object... replacements) {
   }

   public void sendSimpleYesNo(String text) {
   }

   public String numberWithCommas(int number) {
      return "";
   }

   public String getFirstJobStatRequirement(int i) {
      return "";
   }

   public void resetStats() {
   }

   public Optional<Party> getParty() {
      return Optional.empty();
   }

   public boolean isQuestActive(int questId) {
      return false;
   }

   public void sendPinkText(String token) {

   }

   public Pet getPet(int integer) {
      return null;
   }

   public boolean canHoldAll(List<Integer> itemIds) {
      return false;
   }

   public boolean canHoldAll(List<Integer> itemIds, List<Integer> quantities) {
      return false;
   }

   public void disconnect(Boolean aBoolean1, Boolean aBoolean2) {
   }

   public String getMapName() {
      return "";
   }

   public List<Pet> getDriedPets() {
      return Collections.emptyList();
   }
}