package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NpcCommands {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String common_heading = "@"
   String staff_heading = "!"
   String[] levels = ["Common", "Donator", "JrGM", "GM", "SuperGM", "Developer", "Admin"]
   List<Pair<List<String>, List<String>>> commands

   def start() {
      status = -1
      commands = CommandsExecutor.getInstance().getGmCommands()
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            String sendStr = "There are all available commands for you:\r\n\r\n#b"
            for (int i = 0; i <= cm.gmLevel(); i++ ) {
               sendStr += "#L" + i + "#" + levels[i] + "#l\r\n"
            }

            cm.sendSimple(sendStr)
         } else if (status == 1) {
            String[] lvComm, lvDesc
            String lvHead = (selection < 2) ? common_heading : staff_heading

            if (selection > 6) {
               selection = 6
            } else if (selection < 0) {
               selection = 0
            }

            lvComm = commands.get(selection).getLeft()
            lvDesc = commands.get(selection).getRight()

            String sendStr = "The following commands are available for #b" + levels[selection] + "#k:\r\n\r\n"
            for (int i = 0; i < lvComm.size(); i++ ) {
               sendStr += "  #L" + i + "# " + lvHead + lvComm[i] + " - " + lvDesc[i]
               sendStr += "#l\r\n"
            }

            cm.sendPrev(sendStr)
         } else {
            cm.dispose()
         }
      }
   }
}

NpcCommands getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NpcCommands(cm: cm))
   }
   return (NpcCommands) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }