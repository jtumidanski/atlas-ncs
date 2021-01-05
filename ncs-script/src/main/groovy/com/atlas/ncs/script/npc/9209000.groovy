package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9209000 {
   NPCConversationManager cm
   int status = -1
   int selected = -1

   int[] skillBooks = [], masteryBooks = [], table = []

   def start() {
      status = -1
      selected = 0
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
            String greeting = "Hello, I'm #p9209000#, the Skill & Mastery Book announcer! "

            if (cm.isCygnus()) {
               cm.sendOk(greeting + "There are no skill or mastery books available for Cygnus Knights.")
               cm.dispose()
               return
            }

            int jobRank = cm.getJobId() % 10
            if (jobRank < 2) {
               cm.sendOk(greeting + "Keep training yourself until you reach the #r4th job#k of your class. New opportunities for improvement will arrive when you reach that feat!")
               cm.dispose()
               return
            }

            skillBooks = cm.getAvailableSkillBooks()
            masteryBooks = cm.getAvailableMasteryBooks()

            if (skillBooks.length == 0 && masteryBooks.length == 0) {
               cm.sendOk(greeting + "There are no books available to further improve your job skills for now. Either you #bmaxed out everything#k or #byou didn't reach the minimum requisites to use some skill books#k yet.")
               cm.dispose()
            } else if (skillBooks.length > 0 && masteryBooks.length > 0) {
               String sendStr = greeting + "New opportunities for skill improvement have been located for you to improve your skills! Pick a type to take a look onto.\r\n\r\n#b"

               sendStr += "#L1# Skill Book#l\r\n"
               sendStr += "#L2# Mastery Book#l\r\n"

               cm.sendSimple(sendStr)
            } else if (skillBooks.length > 0) {
               selected = 1
               cm.sendNext(greeting + "New opportunities for skill improvement have been located for you to improve your skills! Only skill learns available for now.")
            } else {
               selected = 2
               cm.sendNext(greeting + "New opportunities for skill improvement have been located for you to improve your skills! Only skill upgrades available.")
            }

         } else if (status == 1) {
            String sendStr = "The following books are currently available:\r\n\r\n"
            if (selected == 0) {
               selected = selection
            }

            if (selected == 1) {
               table = skillBooks
               for (int i = 0; i < table.length; i++) {
                  if (table[i] > 0) {
                     int itemId = table[i]
                     sendStr += "  #L" + i + "# #i" + itemId + "#  #t" + itemId + "##l\r\n"
                  } else {
                     int skillId = -table[i]
                     sendStr += "  #L" + i + "# #s" + skillId + "#  #q" + skillId + "##l\r\n"
                  }
               }
            } else {
               table = masteryBooks
               for (int i = 0; i < table.length; i++) {
                  int itemId = table[i]
                  sendStr += "  #L" + i + "# #i" + itemId + "#  #t" + itemId + "##l\r\n"
               }
            }

            cm.sendSimple(sendStr)

         } else if (status == 2) {
            selected = selection

            String sendStr
            if (table[selected] > 0) {
               String[] mobList = cm.getNamesWhoDropsItem(table[selected])

               if (mobList.length == 0) {
                  sendStr = "No mobs drop '#b#t" + table[selected] + "##k'.\r\n\r\n"
               } else {
                  sendStr = "The following mobs drop '#b#t" + table[selected] + "##k':\r\n\r\n"

                  for (int i = 0; i < mobList.length; i++) {
                     sendStr += "  #L" + i + "# " + mobList[i] + "#l\r\n"
                  }

                  sendStr += "\r\n"
               }

            } else {
               sendStr = "\r\n\r\n"
            }
            sendStr += cm.getSkillBookInfo(table[selected])

            cm.sendNext(sendStr)
            cm.dispose()
         }
      }
   }
}

NPC9209000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9209000(cm: cm))
   }
   return (NPC9209000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }