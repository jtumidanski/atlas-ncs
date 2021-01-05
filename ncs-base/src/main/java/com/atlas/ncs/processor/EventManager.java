package com.atlas.ncs.processor;

import java.util.Collections;
import java.util.List;

import com.atlas.ncs.model.Party;
import com.atlas.ncs.model.PartyCharacter;

public class EventManager {
   public Object getProperty(String propertyName) {
      return null;
   }

   public PartyCharacter[] getEligibleParty(Party party) {
      return new PartyCharacter[0];
   }

   public boolean startInstance(int characterId) {
      return false;
   }

   public boolean startInstance(Expedition expedition) {
      return false;
   }

   public boolean startInstance(Party party, int mapId, int something) {
      return false;
   }

   public boolean startInstance(int lobbyId, Party party, int mapId, int something) {
      return false;
   }

   public void setProperty(String propertyName, String value) {
   }

   public void setProperty(String propertyName, int value) {
   }

   public List<EventInstanceManager> getInstances() {
      return Collections.emptyList();
   }

   public byte addGuildToQueue(int guildId, int characterId) {
      return 0;
   }

   public EventInstanceManager getInstance(String name) {
      return null;
   }

   /**
    * Returns true if the guild queue is full.
    * @return
    */
   public boolean isQueueFull() {
      return false;
   }

   /**
    * Gets the size of the guild queue.
    * @return
    */
   public int getQueueSize() {
      return 0;
   }
}
