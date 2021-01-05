package com.atlas.ncs.configuration;

public class Configuration {
   public boolean enableSoloExpeditions() {
      //YamlConfig.config.server.USE_ENABLE_SOLO_EXPEDITIONS
      return false;
   }

   public boolean useCPQ() {
      //YamlConfig.config.server.USE_CPQ
      return false;
   }

   public boolean enableCustomNpcScript() {
      // YamlConfig.config.server.USE_ENABLE_CUSTOM_NPC_SCRIPT
      return false;
   }

   public boolean fastDojo() {
      return false;
   }

   public boolean useRebirthSystem() {
      return false;
   }

   public int weddingBlessingExperience() {
      return 0;
   }

   public boolean weddingBlesserShowFx() {
      return false;
   }
}
