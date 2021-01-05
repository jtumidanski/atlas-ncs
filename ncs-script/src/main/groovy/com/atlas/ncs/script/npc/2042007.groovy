package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2042007 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int rnk = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (status >= 0 && mode == 0) {
            cm.sendOk("2042007_CHAT_LATER")

            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (cm.getMapId() == 980030010) {
            if (status == 0) {
               cm.sendNext("2042007_HOPE_YOU_HAD_FUN")

            } else if (status > 0) {
               cm.warp(980030000, 0)
               cm.dispose()
            }
         } else if (cm.isCPQLoserMap()) {
            if (status == 0) {
               if (cm.getParty() != null) {
                  String shiu = ""
                  if (cm.getFestivalPoints() >= 300) {
                     shiu += "#rA#k"
                     cm.sendOk("Unfortunately, you either drew or lost the battle despite your excellent performance. Victory can be yours next time! \r\n\r\n#bYour result: " + shiu)
                     rnk = 10
                  } else if (cm.getFestivalPoints() >= 100) {
                     shiu += "#rB#k"
                     rnk = 20
                     cm.sendOk("Unfortunately, you either drew or lost the battle, even with your ultimate performance. Just a little bit, and the victory could have been yours! \r\n\r\n#bYour result: " + shiu)
                  } else if (cm.getFestivalPoints() >= 50) {
                     shiu += "#rC#k"
                     rnk = 30
                     cm.sendOk("Unfortunately, you either drew or lost the battle. Victory is for those who strive. I see your efforts, so victory is not far from your reach. Keep it up!\r\n\r\n#bYour result: " + shiu)
                  } else {
                     shiu += "#rD#k"
                     rnk = 40
                     cm.sendOk("Unfortunately, you either equalized or lost the battle, and your performance clearly reflects on it. I expect more from you next time. \r\n\r\n#bYour result: " + shiu)
                  }
               } else {
                  cm.warp(980030000, 0)
                  cm.dispose()
               }
            } else if (status == 1) {
               switch (rnk) {
                  case 10:
                     cm.warp(980030000, 0)
                     cm.gainExp(35000)
                     cm.dispose()
                     break
                  case 20:
                     cm.warp(980030000, 0)
                     cm.gainExp(25000)
                     cm.dispose()
                     break
                  case 30:
                     cm.warp(980030000, 0)
                     cm.gainExp(12500)
                     cm.dispose()
                     break
                  case 40:
                     cm.warp(980030000, 0)
                     cm.gainExp(3500)
                     cm.dispose()
                     break
                  default:
                     cm.warp(980030000, 0)
                     cm.dispose()
                     break
               }
            }
         } else if (cm.isCPQWinnerMap()) {
            if (status == 0) {
               if (cm.getParty() != null) {
                  String shi = ""
                  if (cm.getFestivalPoints() >= 300) {
                     shi += "#rA#k"
                     rnk = 1
                     cm.sendOk("Congratulations on your victory!!! What a performance! The opposite group could not do anything! I hope the same good work next time! \r\n\r\n#bYour result: " + shi)
                  } else if (cm.getFestivalPoints() >= 100) {
                     shi += "#rB#k"
                     rnk = 2
                     cm.sendOk("Congratulations on your victory! That was awesome! You did a good job against the opposing group! Just a little longer, and you'll definitely get an A next time! \r\n\r\n#bYour result: " + shi)
                  } else if (cm.getFestivalPoints() >= 50) {
                     shi += "#rC#k"
                     rnk = 3
                     cm.sendOk("Congratulations on your victory. You did some things here and there, but that can not be considered a good victory. I expect more from you next time. \r\n\r\n#bYour result: " + shi)
                  } else {
                     shi += "#rD#k"
                     rnk = 4
                     cm.sendOk("Congratulations on your victory, though your performance did not quite reflect that. Be more active in your next participation in the Monster Carnival! \r\n\r\n#bYour result: " + shi)
                  }
               } else {
                  cm.warp(980030000, 0)
                  cm.dispose()
               }
            } else if (status == 1) {
               switch (rnk) {
                  case 1:
                     cm.warp(980030000, 0)
                     cm.gainExp(875000)
                     cm.dispose()
                     break
                  case 2:
                     cm.warp(980030000, 0)
                     cm.gainExp(700000)
                     cm.dispose()
                     break
                  case 3:
                     cm.warp(980030000, 0)
                     cm.gainExp(555000)
                     cm.dispose()
                     break
                  case 4:
                     cm.warp(980030000, 0)
                     cm.gainExp(100000)
                     cm.dispose()
                     break
                  default:
                     cm.warp(980030000, 0)
                     cm.dispose()
                     break
               }
            }
         }
      }
   }
}

NPC2042007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2042007(cm: cm))
   }
   return (NPC2042007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }