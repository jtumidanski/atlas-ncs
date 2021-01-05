package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2010000 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int[] eQuestChoices = [4000073, 4000059, 4000060, 4000061, 4000058,
                          4000062, 4000048, 4000049, 4000055, 4000056,
                          4000051, 4000052, 4000050, 4000057, 4000053,
                          4000054, 4000076, 4000078, 4000081, 4000070,
                          4000071, 4000072, 4000069, 4000079, 4000080]
   int[][][] eQuestPrizes = [

         [[2000001, 20],  // Orange Potions
          [2010004, 10],   // Lemons
          [2000003, 15],   // Blue Potions
          [4003001, 15],   // Processed Wood
          [2020001, 15],   // Fried Chickens
          [2030000, 15]],   // Nearest Town Scroll

         [[2000003, 20],   // Blue Potions
          [2000001, 30],   // Orange Potions
          [2010001, 40],   // Meats
          [4003001, 20],   // Processed Wood
          [2040002, 1]],   // 10% Helm Def

         [[2000002, 25],   // White Potions
          [2000006, 10],   // Mana Elixir
          [2022000, 5],   // Pure Water
          [4000030, 15],   // Dragon Skins
          [2040902, 1]],   // 10% Shield Def

         [[2000002, 30],   // White Potions
          [2000006, 15],   // Mana Elixir
          [2020000, 20],   // Salad
          [4003000, 5],   // Screws
          [2041016, 1]],   // 10% Cape Int

         [[2000002, 15],   // White Potions
          [2010004, 15],   // Lemons
          [2000003, 25],   // Blue Potions
          [4003001, 30],   // Processed Wood
          [2040302, 1]],   // 10% Earring Int

         [[2000002, 30],   // White Potions
          [2000006, 15],   // Mana Elixir
          [2020000, 20],   // Salad
          [4003000, 5],   // Screws
          [2040402, 1]],   // 10% Top Def

         [[2000002, 30],   // White Potions
          [2020000, 20],   // Salad
          [2000006, 15],   // Mana Elixir
          [4003000, 5],   // Screws
          [2040402, 1]],   // 10% Top Def

         [[2000006, 25],   // Mana Elixir
          [2020000, 20],   // Salad
          [4020000, 7],   // Garnet Ore
          [4020001, 7],   // Amethyst Ore
          [4020002, 3],   // Aquamarine Ore
          [4020007, 2],   // Diamond Ore
          [2040708, 1]],   // 10% Shoe Speed

         [[2020005, 30],   // Hotdogs
          [2020006, 15],   // Hotdog Supremes
          [2022001, 30],   // Red Bean Soup
          [4003003, 1],   // Fairy Wing
          [2040505, 1]],   // 10% O/All Def

         [[2000006, 25],   // Mana Elixir
          [4020005, 7],   // Sapphire Ore
          [4020003, 7],   // Emerald Ore
          [4020004, 7],   // Opal Ore
          [4020008, 2],   // Black Crystal Ore
          [2040802, 1]],   // 10% Glove Dex

         [[2002004, 15],   // Warrior Potion
          [2002005, 15],   // Sniper Potion
          [2002003, 15],   // Wizard Potion
          [4001005, 1],   // Ancient Scroll
          [2040502, 1]],   // 10% O/All Dex

         [[2000006, 20],   // Mana Elixir
          [4010004, 7],   // Silver Ore
          [4010003, 7],   // Adamantium Ore
          [4010005, 7],   // Orihalcon Ore
          [4003002, 1],   // Piece of Ice
          [2040602, 1]],   // 10% Bottom Def

         [[2000006, 20],   // Mana Elixir
          [4010002, 7],   // Mithril Ore
          [4010001, 7],   // Steel Ore
          [4010000, 7],   // Bronze Ore
          [4010006, 2],   // Gold Ore
          [4003000, 5],   // Screw
          [2040702, 1]],   // 10% Shoe Dex

         [[2000006, 20],   // Mana Elixir
          [4010004, 7],   // Silver Ore
          [4010005, 7],   // Orihalcon Ore
          [4010006, 3],   // Gold Ore
          [4020007, 2],   // Diamond Ore
          [4020008, 2],   // Black Crystal Ore
          [2040705, 1]],   // 10% Shoe Jump

         [[2000006, 30],   // Mana Elixir
          [4020006, 7],   // Topaz Ore
          [4020008, 2],   // Black Crystal Ore
          [4020007, 2],   // Diamond Ore
          [2070010, 1],   // Icicle Stars
          [2040805, 1]],   // 10% Glove Attack

         [[2000006, 30],   // Mana Elixir
          [4020006, 7],   // Topaz Ore
          [4020008.2],   // Black Crystal Ore
          [4020007, 2],   // Diamond Ore
          [2041020, 1]],   // 10% Cape Dex

         [[2000001, 30],   // Orange Potions
          [2000003, 20],   // Blue Potions
          [4003001, 20],   // Processed Wood
          [2010001, 40],   // Meats
          [2040002, 1]],   // 10% Helm Def

         [[2000002, 15],   // White Potions
          [2000003, 25],   // Blue Potions
          [2010004, 15],   // Lemons
          [2050004, 15],   // Divine Elixir
          [4003001, 30],   // Processed Wood
          [2040302, 1]],   // 10% Earring Int

         [[2000006, 25],   // Mana Elixir
          [2020006, 25],   // Hotdog Supreme
          [4010004, 8],   // Silver Ore
          [4010005, 8],   // Orihalcon Ore
          [4010006, 3],   // Gold Ore
          [4020007, 2],   // Diamond Ore
          [4020008, 2],   // Black Crystal Ore
          [2040705, 1]],   // 10% Shoe Jump

         [[2000002, 30],   // White Potions
          [2020000, 20],   // Salad
          [2000006, 15],   // Mana Elixir
          [4003000, 5],   // Screws
          [2041005, 1]],   // 10% Cape Wep Def

         [[2000002, 30],   // White Potions
          [2020000, 20],   // Salad
          [2000006, 15],   // Mana Elixir
          [4003000, 5],   // Screws
          [2041005, 1]],   // 10% Cape Wep Def

         [[2000002, 30],   // White Potions
          [2020000, 20],   // Salad
          [2000006, 15],   // Mana Elixir
          [4003000, 5],   // Screws
          [2041005, 1]],   // 10% Cape Wep Def

         [[2000006, 20],   // Mana Elixir
          [2020005, 30],   // Hotdogs
          [2020006, 15],   // Hotdog Supremes
          [2050004, 20],   // Divine Elixirs
          [4003003, 1],   // Fairy Wing
          [2041002, 1]],   // 10% Cape Mag Def

         [[2000006, 25],   // Mana Elixir
          [2050004, 50],   // Divine Elixir
          [2022001, 35],   // Red Bean Soup
          [4020000, 8],   // Garnet Ore
          [4020001, 8],   // Amethyst Ore
          [4020002, 8],   // Aquamarine Ore
          [4020007, 2],   // Diamond Ore
          [2041023, 1]],   // 10% Cape LUK

         [[2000006, 35],   // Mana Elixir
          [4020006, 9],   // Topaz Ore
          [4010008, 4],   // Black Crystal Ore
          [4020007, 4],   // Diamond Ore
          [2041008, 1]]]   // 10% Cape HP

   int requiredItem = 0
   int lastSelection = 0
   int prizeItem = 0
   int prizeQuantity = 0
   int itemSet
   int[][] reward

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode <= 0) {
         cm.sendOk("2010000_SHOULD_NOT_BE_A_BAD_DEAL")
         cm.dispose()
         return
      }

      status++
      if (status == 0) { // first interaction with NPC
         cm.sendNext("2010000_GOT_A_BIT_OF_TIME")
      } else if (status == 1) {
         cm.sendYesNo("2010000_DEAL_IS_SIMPLE")
      } else if (status == 2) {
         String eQuestChoice = makeChoices(eQuestChoices)
         cm.sendSimple(eQuestChoice)
      } else if (status == 3) {
         lastSelection = selection
         requiredItem = eQuestChoices[selection]
         cm.sendYesNo("2010000_YOU_WANT_TO_TRADE", requiredItem)
      } else if (status == 4) {
         itemSet = (Math.floor(Math.random() * eQuestPrizes[lastSelection].length)).intValue()
         reward = eQuestPrizes[lastSelection]
         prizeItem = reward[itemSet][0]
         prizeQuantity = reward[itemSet][1]
         if (!cm.haveItem(requiredItem, 100)) {
            cm.sendOk("2010000_INVENTORY_FULL", requiredItem)
         } else if (!cm.canHold(prizeItem)) {
            cm.sendOk("2010000_USE_AND_ETC_INVENTORY_FULL")
         } else {
            cm.gainItem(requiredItem, (short) -100)
            cm.gainExp(500 * cm.getExpRate())
            cm.gainItem(prizeItem, (short) prizeQuantity)
            cm.sendOk("2010000_WHAT_DO_YOU_THINK", requiredItem, prizeQuantity, prizeItem)
         }
         cm.dispose()
      }
   }

   static def makeChoices(int[] a) {
      String result = "Ok! First you need to choose the item that you'll trade with. The better the item, the more likely the chance that I'll give you something much nicer in return.\r\n"
      for (int x = 0; x < a.length; x++) {
         result += " #L" + x + "##v" + a[x] + "#  #t" + a[x] + "##l\r\n"
      }
      return result
   }
}

NPC2010000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2010000(cm: cm))
   }
   return (NPC2010000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }