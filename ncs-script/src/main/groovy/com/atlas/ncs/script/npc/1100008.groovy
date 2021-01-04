package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1100008 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String[] menu = ["Ereve"]

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && status == 0) {
            cm.dispose()
            return
         } else if (mode == 0) {
            cm.sendNext("1100008_LET_ME_KNOW")
            cm.dispose()
            return
         }
         status++
         if (status == 0) {
            String display = ""
            for (def i = 0; i < menu.length; i++) {
               display += "\r\n#L" + i + "##b Ereve (1000 mesos)#k"
            }
            cm.sendSimple("This ship will head towards #bEreve#k, an island where you'll find crimson leaves soaking up the sun, the gentle breeze that glides past the stream, and the Empress of Maple Cygnus. If you're interested in joining the Cygnus Knights, Then you should definitely pay a visit here. Are you interested in visiting Ereve?, The Trip will cost you #b1000#k Mesos\r\n" + display)

         } else if (status == 1) {
            if (cm.getMeso() < 1000) {
               cm.sendNext("1100008_NOT_ENOUGH_MESOS")
               cm.dispose()
            } else {
               cm.gainMeso(-1000)
               cm.warp(200090020)
               cm.dispose()
            }
         }
      }
   }
}

NPC1100008 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1100008(cm: cm))
   }
   return (NPC1100008) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }