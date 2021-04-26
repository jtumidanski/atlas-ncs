package com.atlas.ncs.processor;

import java.awt.*;
import java.text.MessageFormat;
import java.util.Collections;
import java.util.List;
import java.util.Locale;
import java.util.Optional;
import java.util.ResourceBundle;
import java.util.concurrent.CompletableFuture;

import com.atlas.ncs.NPCScriptRegistry;
import com.atlas.ncs.configuration.Configuration;
import com.atlas.ncs.event.producer.ChangeMapCommandProducer;
import com.atlas.ncs.event.producer.CharacterEnableActionsProducer;
import com.atlas.ncs.event.producer.CharacterExperienceGainProducer;
import com.atlas.ncs.event.producer.AdjustMesoProducer;
import com.atlas.ncs.event.producer.NpcTalkCommandProducer;
import com.atlas.ncs.event.producer.ServerNoticeProducer;
import com.atlas.ncs.model.Alliance;
import com.atlas.ncs.model.GuildCharacter;
import com.atlas.ncs.model.MapObject;
import com.atlas.ncs.model.Monster;
import com.atlas.ncs.model.NPC;
import com.atlas.ncs.model.Party;
import com.atlas.ncs.model.Pet;
import com.atlas.ncs.model.Portal;

public class NPCConversationManager {
   private final int worldId;

   private final int channelId;

   private final int mapId;

   private final int characterId;

   private final int npcId;

   private final int npcOid;

   private final String scriptName;

   private final boolean itemScript;

   private String text;

   public NPCConversationManager(int worldId, int channelId, int mapId, int characterId, int npc) {
      this(worldId, channelId, mapId, characterId, npc, -1, "", false);
   }

   public NPCConversationManager(int worldId, int channelId, int mapId, int characterId, int npc, String scriptName) {
      this(worldId, channelId, mapId, characterId, npc, -1, scriptName, false);
   }

   public NPCConversationManager(int worldId, int channelId, int mapId, int characterId, int npcId, int oid, String scriptName,
                                 boolean itemScript) {
      this.worldId = worldId;
      this.channelId = channelId;
      this.mapId = mapId;
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
      ResourceBundle bundle = ResourceBundle.getBundle("MessageBundle", Locale.US);
      String message = bundle.getString(token);
      if (replacements == null || replacements.length == 0) {
         return new MessageFormat(message, Locale.US).format(replacements);
      } else {
         return message;
      }
   }

   public void sendSimple(String token, Object... replacements) {
      NpcTalkCommandProducer.sendSimple(characterId, npcId, evaluateToken(token, replacements));
   }

   public void sendNext(String token, Object... replacements) {
      NpcTalkCommandProducer.sendNext(characterId, npcId, evaluateToken(token, replacements));
   }

   public void sendNextSpeaker(String token, byte speaker) {
      NpcTalkCommandProducer.sendNext(characterId, npcId, evaluateToken(token), speaker);
   }

   public void sendNextPrev(String token, Object... replacements) {
      NpcTalkCommandProducer.sendNextPrevious(characterId, npcId, evaluateToken(token, replacements));
   }

   public void sendNextPrevSpeaker(String token, byte speaker) {
      NpcTalkCommandProducer.sendNextPrevious(characterId, npcId, evaluateToken(token), speaker);
   }

   public void sendYesNo(String token, Object... replacements) {
      NpcTalkCommandProducer.sendYesNo(characterId, npcId, evaluateToken(token, replacements));
   }

   /**
    * Warps the set of characters to the identified map and default portal.
    *
    * @param characterIds the characters to warp
    * @param mapId        the map identifier
    */
   public void warp(List<Integer> characterIds, int mapId) {
      characterIds.forEach(id -> warp(id, mapId, 0));
   }

   /**
    * Warps the given character to the identified map and portal.
    *
    * @param characterId the character to warp
    * @param mapId       the map identifier
    * @param portalId    the portal identifier
    */
   public void warp(int characterId, int mapId, int portalId) {
      ChangeMapCommandProducer.changeMap(worldId, channelId, characterId, mapId, portalId);
   }

   /**
    * Warps the current character to the identified map and portal.
    *
    * @param mapId    the map identifier
    * @param portalId the portal identifier
    */
   public void warp(int mapId, int portalId) {
      warp(characterId, mapId, portalId);
   }

   /**
    * Warps the current character to the identified map and portal.
    *
    * @param mapId      the map identifier
    * @param portalName the portal name
    */
   public void warp(int mapId, String portalName) {
      PortalProcessor.getMapPortalByName(mapId, portalName)
            .thenApply(Portal::id)
            .exceptionally(fn -> 0)
            .thenAccept(id -> warp(mapId, id));
   }

   /**
    * Warps the current character to the identified map and default portal.
    *
    * @param mapId the map identifier
    */
   public void warp(int mapId) {
      warp(mapId, 0);
   }

   public void sendPrev(String token, Object... replacements) {
      NpcTalkCommandProducer.sendPrevious(characterId, npcId, evaluateToken(token, replacements));
   }

   public void sendOk(String token, Object... replacements) {
      NpcTalkCommandProducer.sendOk(characterId, npcId, evaluateToken(token, replacements));
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
      AdjustMesoProducer.command(characterId, meso);
   }

   /**
    * Gets the meso amount of the current character.
    *
    * @return the current meso
    */
   public int getMeso() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .meso();
   }

   /**
    * Gets the level of the current character.
    *
    * @return the characters level
    */
   public int getLevel() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .level();
   }

   /**
    * Gets the job identifier for the current character.
    *
    * @return the job identifier
    */
   public int getJobId() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .jobId();
   }

   /**
    * Saves a location of interest for the character.
    *
    * @param type the type of location
    */
   public void saveLocation(String type) {
      CompletableFuture<Point> fromFuture = CharacterProcessor.getCharacter(characterId)
            .thenApply(character -> new Point(character.x(), character.y()))
            .exceptionally(fn -> new Point(0, 0));
      CompletableFuture<List<Portal>> portalsFuture = PortalProcessor.getMapPortals(mapId);

      fromFuture.thenCombine(portalsFuture, this::findClosest)
            .thenApply(Optional::get)
            .thenApply(Portal::id)
            .exceptionally(fn -> 0)
            .thenAccept(id -> CharacterProcessor.saveLocation(characterId, type, mapId, id));
   }

   /**
    * Finds the portal closest to the reference point.
    *
    * @param from    the reference point
    * @param portals the portals to consider
    * @return the closest portal (if one exists).
    */
   protected Optional<Portal> findClosest(Point from, List<Portal> portals) {
      return portals.stream()
            .filter(PortalProcessor::isSpawnPoint)
            .min((o1, o2) -> compareDistanceFromPoint(from, o1, o2));
   }

   /**
    * Compares the distance between the two points, from a point of interest.
    *
    * @param from the point of interest
    * @param o1   the first point
    * @param o2   the second point
    * @return the value 0 if o1 is equal distance to the point of interest as d2; a value less than 0 if o1 is closer
    * than o2 to the point of interest; and a value greater than 0 if o1 is further than d2 from the point of interest.
    */
   protected static int compareDistanceFromPoint(Point from, Portal o1, Portal o2) {
      double o1Distance = new Point(o1.x(), o1.y()).distanceSq(from);
      double o2Distance = new Point(o2.x(), o2.y()).distanceSq(from);
      return Double.compare(o1Distance, o2Distance);
   }

   public int getBuddyListCapacity() {
      return 0;
   }

   public boolean hasItem(int itemId, int quantity) {
      return false;
   }

   public boolean hasItem(int itemId, boolean checkEquipped) {
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

   /**
    * Gets the hair of the current character.
    *
    * @return the hair
    */
   public int getHair() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .hair();
   }

   /**
    * Gets the characters gender.
    *
    * @return 0 if male, 1 if female.
    */
   public int getGender() {
      return characterGender(characterId);
   }

   public void sendStyle(String token, int[] ints) {
   }

   public void setSkin(int skin) {
   }

   public int getMapId() {
      return mapId;
   }

   public boolean isQuestNotStarted(int questId) {
      return false;
   }

   public void forceCompleteQuest(int questId) {
   }

   public void changeJob(int jobId) {
   }

   public void completeQuest(int questId) {
      completeQuest(questId, npcId);
   }

   public void completeQuest(int questId, int npcId) {
   }

   public void sendAcceptDecline(String token, Object... replacements) {
      NpcTalkCommandProducer.sendAcceptDecline(characterId, npcId, evaluateToken(token, replacements));
   }

   public void removeAll(int itemId) {
   }

   /**
    * Returns the mapId for the saved location.
    *
    * @param type the type of location to retrieve.
    * @return the map identifier
    */
   public int getSavedLocation(String type) {
      int location = CharacterProcessor.getSavedLocation(characterId, type);
      CharacterProcessor.saveLocation(characterId, type, 0, 0);
      return location;
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

   /**
    * Shows effect to player
    *
    * @param path
    */
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

   /**
    * Gets the characters current face.
    *
    * @return the face
    */
   public int getFace() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .face();
   }

   /**
    * Returns the mapId for the saved location.
    *
    * @param type the type of location to retrieve.
    * @return the map identifier
    */
   public int peekSavedLocation(String type) {
      return CharacterProcessor.getSavedLocation(characterId, type);
   }

   /**
    * Gets the characters remaining sp total.
    *
    * @return the number of sp remaining
    */
   public int getRemainingSp() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .remainingSp();
   }

   public boolean isGM() {
      return false;
   }

   public int getFestivalPoints() {
      return 0;
   }

   public int getMapMonsterCount() {
      return getMapMonsterCount(getMapId());
   }

   public int getMapMonsterCount(int mapId) {
      return 0;
   }

   public int getMapMonsterCount(int mapId, int area) {
      return 0;
   }

   public int getJobNiche() {
      return 0;
   }

   public double getExpRate() {
      return 1.0;
   }

   public double getMesoRate() {
      return 1;
   }

   public void sendGetNumber(String token, Object... replacements) {
      sendGetNumber(token, 1, 1, 100, replacements);
   }

   public void sendGetNumber(String token, int def, int min, int max, Object... replacements) {
   }

   public void sendSimpleYesNo(String text) {
      NpcTalkCommandProducer.sendYesNo(characterId, npcId, text);
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

   public void sendPinkText(String token, Object... replacements) {
      characterSendPinkText(characterId, token, replacements);
   }

   public void characterSendPinkText(int characterId, String token, Object... replacements) {
      ServerNoticeProducer.sendPinkText(characterId, evaluateToken(token, replacements));
   }

   public Pet getPet(int index) {
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

   public int getHallOfFameMapId(int jobId) {
      //      if (isCygnus(jobId)) {
      //         return 130000100;
      //      } else if (isAran(jobId)) {
      //         return 140010110;
      //      } else {
      //         if (job.isA(MapleJob.WARRIOR)) {
      //            return 102000004;
      //         } else if (job.isA(MapleJob.MAGICIAN)) {
      //            return 101000004;
      //         } else if (job.isA(MapleJob.BOWMAN)) {
      //            return 100000204;
      //         } else if (job.isA(MapleJob.THIEF)) {
      //            return 103000008;
      //         } else if (job.isA(MapleJob.PIRATE)) {
      //            return 120000105;
      //         } else {
      //            return 130000110;   // beginner explorers are allotted with the Cygnus, available map lul
      //         }
      //      }
      return 0;
   }

   public boolean canSpawnPlayerNpc(int mapId) {
      return false;
   }

   public boolean spawnPlayerNPC(int mapId) {
      return false;
   }

   public void cosmeticExistsAndIsNotEquipped(int itemId) {

   }

   public void isRecvPartySearchInviteEnabled() {
   }

   public Optional<EventManager> getEventManager(String eventName) {
      return Optional.empty();
   }

   public boolean isUsingOldPqNpcStyle() {
      return false;
   }

   public boolean isPartyLeader() {
      return false;
   }

   public boolean toggleRecvPartySearchInvite() {
      return false;
   }

   public EventInstanceManager getEventInstance() {
      return null;
   }

   public boolean isEventLeader() {
      return false;
   }

   /**
    * Counts the characters in the given map.
    *
    * @param mapId the map identifier
    * @return the total number of characters
    */
   public int countCharactersInMap(int mapId) {
      return MapProcessor
            .countCharactersInMap(worldId, channelId, mapId)
            .join();
   }

   public void evolvePet(byte petIndex, int newPetId) {
   }

   public void removeFromSlot(String type, short slot, short quantity, boolean fromDrop) {

   }

   public Optional<Monster> getMonster(int monsterId) {
      return Optional.empty();
   }

   public void spawnMonsterOnGroundBelow(int mapId, Monster monster, int x, int y) {
   }

   public void spawnMonsterOnGroundBelow(int monsterId, int x, int y) {
      getMonster(monsterId).ifPresent(monster -> spawnMonsterOnGroundBelow(getMapId(), monster, x, y));
   }

   public void spawnMonsterOnGroundBelow(Monster monster, int x, int y) {
      spawnMonsterOnGroundBelow(getMapId(), monster, x, y);
   }

   /**
    * The current characters gains the given amount of experience.
    *
    * @param amount the amount to gain
    */
   public void gainExp(double amount) {
      CharacterExperienceGainProducer.gainExperience(characterId, (int) Math.round(Math.floor(amount)));
   }

   public int getItemQuantity(int itemId) {
      return 0;
   }

   public String getText() {
      return text;
   }

   public void setText(String text) {
      this.text = text;
   }

   public boolean start_PyramidSubway(Integer integer) {
      return false;
   }

   public boolean bonus_PyramidSubway(Integer integer) {
      return false;
   }

   public String getQuestCustomDataOrDefault(int characterId, int questId, String def) {
      return def;
   }

   public int countFreeInventorySlot(String type) {
      return 0;
   }

   public String getQuestProgress(int questId) {
      return "";
   }

   public int getQuestProgressInt(int questId) {
      return 0;
   }

   public int getQuestProgressInt(int questId, int infoNumber) {
      return 0;
   }

   public int getNpcObjectId() {
      return npcOid;
   }

   public void setQuestProgress(int questId, int infoNumber, int progress) {
   }

   public void setQuestProgress(int questId, int progress) {
   }

   public void setQuestProgress(int questId, String progress) {
   }

   public void teachSkill(int skillId, byte level, byte masterLevel, long expiration) {
   }

   public void unlockUI() {
   }

   public void sendShop(int shopId) {

   }

   public void useItem(int itemId) {

   }

   public boolean isCygnus() {
      return false;
   }

   public int getJobBranch(int jobId) {
      return 0;
   }

   /**
    * Gets the current characters name.
    *
    * @return the name
    */
   public String getCharacterName() {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .name();
   }

   public void destroyNPC(int npcId) {

   }

   public MapObject getMapObject(int id) {
      return null;
   }

   public Optional<NPC> getNpcById(int npcId) {
      return Optional.empty();
   }

   public void spawnNpc(int npcId, int x, int y) {

   }

   public boolean containsAreaInfo(short area, String info) {
      return false;
   }

   public void updateAreaInfo(short area, String info) {
   }

   public void spawnGuide() {
   }

   public int getGuildId() {
      return 0;
   }

   public int getGuildRank() {
      return 0;
   }

   public void genericGuildMessage(int messageId) {
   }

   public void disbandGuild() {
   }

   public void increaseGuildCapacity() {
   }

   public int getGuildCapacity() {
      return 0;
   }

   public String getIncreaseGuildCost(int capacity) {
      return "";
   }

   public int getAllianceId() {
      return 0;
   }

   public Optional<GuildCharacter> getGuildCharacter() {
      return Optional.empty();
   }

   public int getAllianceRank() {
      return 0;
   }

   public int getAllianceCapacity() {
      return 0;
   }

   public void upgradeAlliance() {
   }

   public boolean hasGuild() {
      return false;
   }

   public void disbandAlliance(int allianceId) {

   }

   public boolean canBeUsedAllianceName(String allianceName) {
      return false;
   }

   public Optional<Alliance> createAlliance(String allianceName) {
      return Optional.empty();
   }

   /**
    * Plays sound for all players in map.
    *
    * @param path
    */
   public void playSoundInMap(String path) {
   }

   /**
    * Plays sound for player.
    *
    * @param path
    */
   public void playSound(String path) {
   }

   public int getMapItemCount() {
      return 0;
   }

   /**
    * Hits reactor in map by name.
    *
    * @param name
    * @param hits
    */
   public void forceHitReactor(String name, byte hits) {
   }

   /**
    * Gets reactor state in map by name.
    *
    * @param name
    * @return
    */
   public int getReactorState(String name) {
      return 0;
   }

   public int gmLevel() {
      return 0;
   }

   public int partyMembersInMap() {
      return 0;
   }

   public void sendPopUp(String token) {
   }

   public boolean haveItemEquipped(int itemId) {
      return false;
   }

   public int getSkillLevel(int skillId) {
      return 0;
   }

   public void openShopNPC(int shopId) {
   }

   /**
    * Kills all monsters in map
    */
   public void killAllMonsters() {
   }

   /**
    * Set state of reactor in map
    */
   public void setReactorState() {
   }

   /**
    * Changes music for character.
    *
    * @param path
    */
   public void changeMusic(String path) {

   }

   /**
    * Changes music for map.
    *
    * @param path
    */
   public void changeMusicInMap(String path) {

   }

   public void openUI(Byte ui) {
   }

   public void showInfoText(String text) {

   }

   /**
    * Set party quest for character.
    *
    * @param o
    */
   public void setPartyQuest(Object o) {
   }

   public int itemQuantity(int itemId) {
      return 0;
   }

   public void removeNpc(Integer... npcIds) {
   }

   public void forceStartReactor(Integer... reactorIds) {
   }

   public void setCustomData(int characterId, int questId, String customData) {
   }

   public boolean hasMerchant() {
      return false;
   }

   public boolean hasMerchantItems() {
      return false;
   }

   public void showFredrick() {
   }

   public void gainGP(int guildId, int amount) {

   }

   public void displayGuildRanks() {
   }

   /**
    * Sends a pink text message to all characters in the map.
    *
    * @param token        the token of the message to send
    * @param replacements any replacement values in the tokenized message.
    */
   public void sendPinkTextToMap(String token, Object... replacements) {
      MapProcessor.getCharacterIdsInMap(worldId, channelId, mapId)
            .thenAccept(ids -> ids.forEach(id -> characterSendPinkText(id, token, replacements)));
   }

   public void levelUp(boolean takeExperience) {
   }

   /**
    * Dictates whether the current character is male (or not).
    *
    * @return true if the character is male
    */
   public boolean isMale() {
      return getGender() == 0;
   }

   public int[] getAvailableSkillBooks() {
      return new int[0];
   }

   public int[] getAvailableMasteryBooks() {
      return new int[0];
   }

   public String getSkillBookInfo(int skillId) {
      return "";
   }

   public String[] getNamesWhoDropsItem(int itemId) {
      return new String[0];
   }

   public int getQuestStatus(int questId) {
      return 0;
   }

   public boolean characterHasItem(int characterId, int itemId) {
      return false;
   }

   public boolean characterHasItem(int characterId, int itemId, boolean checkEquipped) {
      return false;
   }

   /**
    * Returns the gender of the given character.
    *
    * @param characterId the character identifier
    * @return 0 if male, 1 if female
    */
   public int characterGender(int characterId) {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .gender();
   }

   public boolean characterHasItemEquipped(int characterId, int itemId) {
      return false;
   }

   public int getPartnerId() {
      return 0;
   }

   public boolean characterCanHold(int characterId, int itemId) {
      return characterCanHold(characterId, itemId, 1);
   }

   public boolean characterCanHold(int characterId, int itemId, int amount) {
      return false;
   }

   public boolean hasPartner() {
      return false;
   }

   public boolean partnerInMap() {
      return false;
   }

   public void characterGainExp(int characterId, int amount) {

   }

   public void characterGainItem(int characterId, int itemId, short amount) {

   }

   /**
    * Sends a light blue text message to all characters in the map.
    *
    * @param token        the token of the message to send
    * @param replacements any replacement values in the tokenized message
    */
   public void sendBlueTextToMap(String token, Object... replacements) {
      MapProcessor.getCharacterIdsInMap(worldId, channelId, mapId)
            .thenAccept(ids -> ids.forEach(id -> characterSendBlueText(id, token, replacements)));
   }

   public void characterSetMarriageItemId(int characterId, int itemId) {
   }

   public void characterNpcTalk(int characterId, int npcId, String message) {

   }

   /**
    * Sends a light blue text message to the given character.
    *
    * @param characterId  the character identifier
    * @param token        the token of the message to send
    * @param replacements any replacement values in the tokenized message
    */
   public void characterSendBlueText(int characterId, String token, Object... replacements) {
      ServerNoticeProducer.sendLightBlueText(characterId, evaluateToken(token, replacements));
   }

   public boolean isMarried() {
      return false;
   }

   /**
    * Gets the map the given character is in.
    *
    * @param characterId the character identifier
    * @return the map identifier
    */
   public int characterGetMap(int characterId) {
      return CharacterProcessor.getCharacter(characterId)
            .join()
            .mapId();
   }

   public void openNpc(int npcId) {

   }

   public void openNpc(int npcId, String script) {

   }

   public boolean isGuildLeader() {
      return false;
   }

   /**
    * Gets the channel identifier the character is in.
    *
    * @return the channel identifier
    */
   public int getChannelId() {
      return channelId;
   }

   public void endExpedition(Expedition mapleExpedition) {
   }

   public Expedition getExpedition(ExpeditionType expeditionType) {
      return null;
   }

   public int createExpedition(ExpeditionType expeditionType) {
      return 0;
   }

   public void removeNPC(int npcId) {

   }

   /**
    * Sends notice to current character
    *
    * @param message the message to send
    */
   public void sendNotice(String message) {
      ServerNoticeProducer.sendNotice(characterId, message);
   }

   public boolean isLeaderExpedition(ExpeditionType expeditionType) {
      return false;
   }

   public String getExpeditionMemberNames(ExpeditionType expeditionType) {
      return null;
   }

   public void setItemExpiration(String inventoryType, int slot, long expiration) {
   }

   public void sendDirectionInfo(int integer1, int integer2) {
      throw new UnsupportedOperationException();
   }

   public void sendDirectionInfo(String path, int integer1, int integer2, int integer3, int integer4, int integer5) {
      throw new UnsupportedOperationException();
   }

   public void updateInfo(String s1, String s2) {
      throw new UnsupportedOperationException();
   }

   public Configuration getConfiguration() {
      return null;
   }

   public boolean isCPQLoserMap() {
      return false;
   }

   public boolean isCPQWinnerMap() {
      return false;
   }

   public boolean sendCPQMapLists() {
      return false;
   }

   public boolean fieldTaken(int fieldId) {
      return false;
   }

   public boolean fieldLobbied(int fieldId) {
      return false;
   }

   public void challengeParty(int fieldId) {
   }

   public void cpqLobby(int fieldId) {
   }

   public boolean sendCPQMapLists2() {
      return false;
   }

   public boolean fieldTaken2(int fieldId) {
      return false;
   }

   public boolean fieldLobbied2(int fieldId) {
      return false;
   }

   public void challengeParty2(int fieldId) {
   }

   public void cpqLobby2(int fieldId) {
   }

   public void sendDefault() {
   }

   public void setDojoStage(int stage) {

   }

   public boolean hasFinishedDojoTutorial() {
      return false;
   }

   public int getDojoStage() {
      return 0;
   }

   public void resetDojoEnergy() {

   }

   public void resetPartyDojoEnergy() {
   }

   public int getDojoPoints() {
      return 0;
   }

   public void setDojoPoints(int amount) {
   }

   public int getVanquisherStage() {
      return 0;
   }

   public int getVanquisherKills() {
      return 0;
   }

   public void setVanquisherStage(int stage) {

   }

   public void setVanquisherKills(int amount) {
   }

   /**
    * Reset dojo map for channel.
    *
    * @param mapId
    */
   public void resetDojoMap(int mapId) {

   }

   /**
    * Returns if characters are in the supplied map.
    *
    * @param mapId the map identifier
    * @return true if characters are present
    */
   public boolean hasCharactersInMap(int mapId) {
      return countCharactersInMap(mapId) > 0;
   }

   public void clearMapObjects() {
   }

   /**
    * Gives each member of the party experience.
    *
    * @param amount       the amount of experience
    * @param characterIds the characters to reward
    */
   public void givePartyExp(int amount, List<Integer> characterIds) {
      characterIds.forEach(id -> gainExp(amount));
   }

   /**
    * Removes chalkboard for player
    */
   public void removeChalkboard() {
   }

   public void openRPSNpc() {
      //PacketCreator.announce(cm.getClient(), new OpenRPSNPC())
   }

   public int countRebirths() {
      return 0;
   }

   public void executeRebirth() {
   }

   public void sendDimensionalMirror(String token) {
   }

   public Point characterPosition(int characterId) {
      return null;
   }

   public void delayedHitReactor(int reactorId, int delay) {

   }

   public void setPortalState(String portalName, boolean state) {

   }

   public int countItem(int itemId) {
      return 0;
   }

   public Rectangle getMapArea(int areaId) {
      return null;
   }

   /**
    * Dictates whether or not the character is alive.
    *
    * @param characterId the character identifier
    * @return true if the character is alive
    */
   public boolean characterIsAlive(int characterId) {
      return CharacterProcessor
            .getCharacter(characterId)
            .join()
            .hp() > 0;
   }

   /**
    * Toggle drops in map.
    */
   public void toggleDrops() {

   }

   public boolean isAllReactorState(int reactorId, int state) {
      return false;
   }

   public Point getMapPortalPosition(String portalName) {
      return null;
   }

   public void changeCharacterName(String newName) {

   }

   public boolean canCreateChar(String proposedName) {
      return false;
   }
}