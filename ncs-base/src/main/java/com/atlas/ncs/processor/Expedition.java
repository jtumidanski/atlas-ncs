package com.atlas.ncs.processor;

import java.util.Collections;
import java.util.List;

import com.atlas.ncs.model.ExpeditionCharacter;

public class Expedition {
   public boolean isLeader(int characterId) {
      return false;
   }

   public boolean isInProgress() {
      return false;
   }

   public boolean isRegistering() {
      return false;
   }

   public boolean contains(int characterId) {
      return false;
   }

   public String getLeaderName() {
      return null;
   }

   public String addMember(int characterId) {
      return null;
   }

   public List<ExpeditionCharacter> getMemberList() {
      return Collections.emptyList();
   }

   public void ban(ExpeditionCharacter character) {

   }

   public int addMemberInt(int characterId) {
      return 0;
   }

   public String getProperty(String propertyName) {
      return null;
   }

   public void setProperty(String propertyName, String value) {

   }

   public void warpTeam(int mapId) {

   }
}
