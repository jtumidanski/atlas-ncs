package com.atlas.ncs;

import com.atlas.ncs.model.ScriptedItem;
import com.atlas.ncs.processor.NPCConversationManager;

import javax.script.Invocable;
import javax.script.ScriptEngine;
import javax.script.ScriptEngineFactory;
import javax.script.ScriptEngineManager;
import javax.script.ScriptException;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class NPCScriptRegistry {
   private static final Object lock = new Object();

   private static volatile NPCScriptRegistry instance;

   private final Map<Integer, NPCConversationManager> cms = new HashMap<>();

   private final Map<Integer, ScriptEngine> scripts = new HashMap<>();

   private final ScriptEngineFactory sef;

   public static NPCScriptRegistry getInstance() {
      NPCScriptRegistry result = instance;
      if (result == null) {
         synchronized (lock) {
            result = instance;
            if (result == null) {
               result = new NPCScriptRegistry();
               instance = result;
            }
         }
      }
      return result;
   }

   private NPCScriptRegistry() {
      sef = new ScriptEngineManager().getEngineByName("groovy").getFactory();
   }

   public boolean isNpcScriptAvailable(int characterId, String fileName) {
      ScriptEngine iv = null;
      if (fileName != null) {
         iv = getScriptEngine("npc/" + fileName, characterId);
      }
      return iv != null;
   }

   protected ScriptEngine getScriptEngine(String path, int characterId) {
      String cachePath = "/service/script/" + path + ".groovy";
      return getScriptEngine(cachePath);
   }

   protected ScriptEngine getScriptEngine(String path) {
      ScriptEngine engine = sef.getScriptEngine();
      engine = evalPrerequisites(engine, getPrerequisites());
      return eval(engine, path);
   }

   protected String[] getPrerequisites() {
      return new String[0];
   }

   protected ScriptEngine evalPrerequisites(ScriptEngine engine, String... paths) {
      ScriptEngine primedEngine = engine;
      for (String path : paths) {
         primedEngine = eval(primedEngine, path);
         if (primedEngine == null) {
            return null;
         }
      }
      return primedEngine;
   }

   protected ScriptEngine eval(ScriptEngine engine, String path) {
      path = "script/src/main/groovy/" + path;
      File scriptFile = null;
      if (new File(path + ".groovy").exists()) {
         scriptFile = new File(path + ".groovy");
      }
      if (scriptFile == null) {
         return null;
      }

      try (FileReader fr = new FileReader(scriptFile)) {
         engine.eval(fr);
      } catch (final ScriptException | IOException t) {
         return null;
      }
      return engine;
   }

   public boolean start(int characterId, int npc) {
      return start(characterId, npc, -1);
   }

   public boolean start(int characterId, int npc, int oid) {
      return start(characterId, npc, oid, null);
   }

   public boolean start(int characterId, int npc, String fileName) {
      return start(characterId, npc, -1, fileName);
   }

   public boolean start(int characterId, int npc, int oid, String fileName) {
      return start(characterId, npc, oid, fileName, false, "cm");
   }

   public boolean start(int characterId, ScriptedItem scriptItem) {
      return start(characterId, scriptItem.npcId(), -1, scriptItem.script(), true, "im");
   }

   public void start(String filename, int characterId, int npc, List<Integer> partyCharacters) {
      try {
         NPCConversationManager cm = new NPCConversationManager(characterId, npc);
         cm.dispose();
         if (cms.containsKey(characterId)) {
            return;
         }
         cms.put(characterId, cm);
         ScriptEngine iv = getScriptEngine("npc/" + filename, characterId);

         if (iv == null) {
            cm.dispose();
            return;
         }
         iv.put("cm", cm);
         scripts.put(characterId, iv);
         try {
            ((Invocable) iv).invokeFunction("start", partyCharacters);
         } catch (final NoSuchMethodException e) {
            try {
               ((Invocable) iv).invokeFunction("start", partyCharacters);
            } catch (final NoSuchMethodException e1) {
               e1.printStackTrace();
            }
         }

      } catch (final Exception ute) {
         System.out.printf("NPC [%d]%n", npc);
         dispose(characterId);
      }
   }

   private boolean start(int characterId, int npc, int oid, String fileName, boolean itemScript, String engineName) {
      try {
         NPCConversationManager cm = new NPCConversationManager(characterId, npc, oid, fileName, itemScript);
         if (cms.containsKey(characterId)) {
            dispose(characterId);
         }
//         if (c.canClickNPC()) {
         cms.put(characterId, cm);
         ScriptEngine iv = null;
         if (!itemScript) {
            if (fileName != null) {
               iv = getScriptEngine("npc/" + fileName, characterId);
            }
         } else {
            if (fileName != null) {
               iv = getScriptEngine("item/" + fileName, characterId);
            }
         }
         if (iv == null) {
            iv = getScriptEngine("npc/" + npc, characterId);
//               cm.resetItemScript();
         }
         if (iv == null) {
            dispose(characterId);
            return false;
         }
         iv.put(engineName, cm);
         scripts.put(characterId, iv);
//            c.setClickedNPC();
         try {
            ((Invocable) iv).invokeFunction("start");
         } catch (final NoSuchMethodException e) {
            e.printStackTrace();
            try {
               ((Invocable) iv).invokeFunction("start", characterId);
            } catch (final NoSuchMethodException e1) {
               e1.printStackTrace();
            }
         }
//         } else {
//            CharacterEnableActionsProducer.enableActions(characterId);
//         }
         return true;
      } catch (final Exception ute) {
         System.out.printf("NPC [%d]%n", npc);
         dispose(characterId);

         return false;
      }
   }

   public void action(int characterId, byte mode, byte type, int selection) {
      ScriptEngine iv = scripts.get(characterId);
      if (iv != null) {
         try {
//            c.setClickedNPC();
            ((Invocable) iv).invokeFunction("action", mode, type, selection);
         } catch (ScriptException | NoSuchMethodException t) {
            if (getCM(characterId) != null) {
               System.out.printf("NPC [%d]%n", getCM(characterId).getNpcId());
            }
            dispose(characterId);
         }
      }
   }

   public void dispose(NPCConversationManager cm) {
//      MapleClient c = cm.getClient();
//      c.getPlayer().setCS(false);
//      c.getPlayer().setNpcCoolDown(System.currentTimeMillis());
      cms.remove(cm.getCharacterId());
      scripts.remove(cm.getCharacterId());
//      c.getPlayer().flushDelayedUpdateQuests();
   }

   public void dispose(int characterId) {
      NPCConversationManager cm = cms.get(characterId);
      if (cm != null) {
         dispose(cm);
      }
   }

   public NPCConversationManager getCM(int characterId) {
      return cms.get(characterId);
   }

}
