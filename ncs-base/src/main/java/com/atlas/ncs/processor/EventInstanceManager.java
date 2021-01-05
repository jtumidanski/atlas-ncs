package com.atlas.ncs.processor;

import java.util.Collections;
import java.util.List;

public class EventInstanceManager {
   public boolean giveEventReward(int characterId) {
      return false;
   }

   public boolean giveEventReward(int characterId, int quantity) {
      return false;
   }

   public void startEventTimer(int time) {
   }

   public void warpEventTeam(int mapId) {
   }

   public void clearPQ() {
   }

   public void setProperty(String propertyName, String value) {

   }

   public void setProperty(String propertyName, String value, boolean previous) {

   }

   public void setProperty(String propertyName, int value) {

   }

   public void showClearEffect() {
      showClearEffect(false);
   }

   public void showClearEffect(boolean hasGate) {
   }

   public void showClearEffect(int mapId) {
      showClearEffect(mapId, false);
   }

   public void showClearEffect(int mapId, boolean hasGate) {
   }

   public void showClearEffect(int mapId, String mapObj, int newState) {

   }

   public void giveEventPlayersStageReward(int stage) {
   }

   public boolean isEventCleared() {
      return false;
   }

   public int getIntProperty(String propertyName) {
      return 0;
   }

   public void warpEventTeamToMapSpawnPoint(int mapId, int portalId) {
   }

   public void giveEventPlayersExp(int amount) {
   }

   public String getProperty(String propertyName) {
      return "";
   }

   public void showWrongEffect() {
   }

   public void setIntProperty(String propertyName, int value) {
   }

   public void gridInsert(int characterId, int integer) {
   }

   public int gridCheck(int characterId) {
      return 0;
   }

   public void linkToNextStage(int thisStage, String eventFamily, int thisMapId) {
   }

   public boolean isEventLeader(int characterId) {
      return false;
   }

   public List<Integer> getClearStageBonus(int stage) {
      return Collections.emptyList();
   }

   public void giveEventPlayersMeso(int amount) {
   }

   public boolean isEventTeamTogether() {
      return false;
   }

   public void restartEventTimer(int time) {
   }

   public void registerPlayer(int characterId) {
   }

   public int gridSize() {
      return 0;
   }

   public void registerParty(int characterId) {

   }

   public void gridClear() {
   }

   public List<Integer> getCharacterIds() {
      return Collections.emptyList();
   }

   public void linkPortalToScript(int stage, String portalName, String scriptName, int mapId) {

   }

   public long getTimeLeft() {
      return 0;
   }
}
