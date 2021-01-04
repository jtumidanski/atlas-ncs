package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9000017 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int selectedType = -1
   int selectedItem = -1
   int item
   int[] mats
   int[] matQty
   int cost
   int qty
   boolean last_use //last item is a use item

   def start() {
      cm.getPlayer().setCS(true)
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.sendOk("9000017_OH_OK")

         cm.dispose()
         return
      }

      if (status == 0) {
         if (!YamlConfig.config.server.USE_ENABLE_CUSTOM_NPC_SCRIPT) {
            cm.sendOk("9000017_HELLO", cm.getNpcId())

            cm.dispose()
            return
         }

         String selStr = "Hey traveler! Come, come closer... We offer a #bhuge opportunity of business#k to you. If you want to know what it is, keep listening..."
         cm.sendNext(selStr)
      } else if (status == 1) {
         String selStr = "We've got here the knowledge to synthesize the mighty #b#t2049100##k! Of course, making one is not an easy task... But worry not! Just gather some material to me and a fee of #b1,200,000 mesos#k for our services to #bobtain it#k. You still want to do it?"
         cm.sendYesNo(selStr)
      } else if (status == 2) {
         //selectedItem = selection;
         selectedItem = 0

         int[] itemSet = [2049100, 7777777]
         int[][] matSet = [[4031203, 4001356, 4000136, 4000082, 4001126, 4080100, 4000021, 4003005]]
         int[][] matQtySet = [[100, 60, 40, 80, 10, 8, 200, 120]]
         int[] costSet = [1200000, 7777777]
         item = itemSet[selectedItem]
         mats = matSet[selectedItem]
         matQty = matQtySet[selectedItem]
         cost = costSet[selectedItem]

         String prompt = "So, you want us to make some #t" + item + "#? In that case, how many do you want us to make?"
         cm.sendGetNumber(prompt, 1, 1, 100)
      } else if (status == 3) {
         qty = (selection > 0) ? selection : (selection < 0 ? -selection : 1)
         last_use = false

         String prompt = "You want us to make "
         if (qty == 1) {
            prompt += "a #t" + item + "#?"
         } else {
            prompt += qty + " #t" + item + "#?"
         }

         prompt += " In that case, we're going to need specific items from you in order to make it. Make sure you have room in your inventory, though!#b"

         for (int i = 0; i < mats.length; i++) {
            prompt += "\r\n#i" + mats[i] + "# " + matQty[i] * qty + " #t" + mats[i] + "#"
         }

         if (cost > 0) {
            prompt += "\r\n#i4031138# " + cost * qty + " meso"
         }
         cm.sendYesNo(prompt)
      } else if (status == 4) {
         boolean complete = true

         if (cm.getMeso() < cost * qty) {
            cm.sendOk("9000017_NOT_DOING_YOU_A_FAVOR")

         } else if (!cm.canHold(item, qty)) {
            cm.sendOk("9000017_NEED_INVENTORY_SPACE")

         } else {
            for (int i = 0; complete && i < mats.length; i++) {
               if (matQty[i] * qty == 1) {
                  complete = cm.haveItem(mats[i])
               } else {
                  complete = cm.haveItem(mats[i], matQty[i] * qty)
               }
            }

            if (!complete) {
               cm.sendOk("9000017_KIDDING_RIGHT")

            } else {
               for (int i = 0; i < mats.length; i++) {
                  cm.gainItem(mats[i], (short) (-matQty[i] * qty))
               }
               cm.gainMeso(-cost * qty)
               cm.gainItem(item, (short) qty)
               cm.sendOk("9000017_CANNOT_BELIEVE_IT_WORKED")

            }
         }
         cm.dispose()
      }
   }
}

NPC9000017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9000017(cm: cm))
   }
   return (NPC9000017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }