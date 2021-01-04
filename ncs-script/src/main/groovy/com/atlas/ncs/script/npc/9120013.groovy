package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9120013 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   List<String> questions
   List<List<String>> answers
   List<Integer> correctAnswer
   int questionNum

   def start() {
      status = -1
      questions = ["Which of these items does the Flaming Raccoon NOT drop?", "Which NPC is responsible for transporting travellers from Kerning City to Zipangu, and back?", "Which of the items sold at the Mushroom Shrine increases your attack power?", "Which of these items do the Extras NOT drop?", "Which of these items DO NOT exist??", "What's the name of the vegetable store owner in Showa Town?", "Which of these items DO exist?", "What is the name of the strongest boss in the Mushroom Shrine?", "Which one of these items has a mis-matched class or level description?", "Which of these noodles are NOT being sold by Robo at the Mushroom Shrine?", "Which of these NPCs do NOT stand in front of Showa Movie Theater?"]
      answers = [["Raccoon Firewood", "Solid Horn", "Red Brick"], ["Peli", "Spinel", "Poli"], ["Takoyaki", "Yakisoba", "Tempura"], ["Extra A's Badge", "Extra B's Corset", "Extra C's Necklace"], ["Frozen Tuna", "Fan", "Fly Swatter"], ["Sami", "Kami", "Umi"], ["Cloud Fox's Tooth", "Ghost's Bouquet", "Dark Cloud Fox's Tail"], ["Black Crow", "Blue Mushmom", "Himegami"], ["Bamboo Spear - Warrior-only Weapon", "Pico-Pico Hammer - One-handed Sword", "Mystic Cane - Level 51 equip"], ["Kinoko Ramen (Pig Skull)", "Kinoko Ramen (Salt)", "Mushroom Miso Ramen"], ["Skye", "Furano", "Shinta"]]
      correctAnswer = [1, 1, 0, 1, 2, 2, 2, 0, 0, 2, 2]
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0 && mode == 1) {
            if (cm.isQuestStarted(8012) && !cm.haveItem(4031064)) { //quest in progress
               cm.sendYesNo("9120013_GET_THEM_ALL")

            } else { //quest not started or already completed
               //cm.sendOk("9120013_MEOW")

               cm.dispose()
            }
         } else if (status == 1 && mode == 1) {
            boolean hasChicken = true
            if (!cm.haveItem(2020001, 300)) {
               hasChicken = false
            }
            if (!hasChicken) {
               cm.sendOk("9120013_NOT_ALL_OF_US")

               cm.dispose()
            } else {
               cm.gainItem(2020001, (short) -300)
               cm.sendNext("9120013_GOOD_JOB")

            }
         } else if (status == 7 && mode == 1) { //2-6 are the questions
            if (selection != correctAnswer.pop()) {
               cm.sendNext("9120013_ALL_HUMANS_MAKE_MISTAKES")

               cm.dispose()
            } else {
               cm.sendNext("9120013_ALL_QUESTIONS_RIGHT")

            }
         } else if (status == 8 && mode == 1) { //gain marble
            cm.gainItem(4031064, (short) 1)
            cm.sendOk("9120013_CAN_LEAVE_NOW")

            cm.dispose()
         } else if (status >= 2 && status <= 6 && mode == 1) {//questions
            boolean cont = true
            if (status > 2) {
               if (selection != correctAnswer.pop()) {
                  cm.sendNext("9120013_ALL_HUMANS_MAKE_MISTAKES")

                  cm.dispose()
                  cont = false
               }
            }
            if (cont) {
               questionNum = Math.floor(Math.random() * questions.size()).intValue()
               if (questionNum != (questions.size() - 1)) {
                  String temp
                  temp = questions[questionNum]
                  questions.set(questionNum, questions.get(questions.size() - 1))
                  questions.set(questions.size() - 1, temp)
                  List<String> temp2 = answers.get(questionNum)
                  answers.set(questionNum, answers.get(questions.size() - 1))
                  answers.set(questions.size() - 1, temp2)
                  int temp3 = correctAnswer.get(questionNum)
                  correctAnswer.set(questionNum, correctAnswer.get(questions.size() - 1))
                  correctAnswer.set(questions.size() - 1, temp3)
               }
               String question = questions.pop()
               String[] answer = answers.pop()
               String prompt = "Question no." + (status - 1) + ": " + question
               for (int i = 0; i < answer.length; i++) {
                  prompt += "\r\n#b#L" + i + "#" + answer[i] + "#l#k"
               }
               cm.sendSimple(prompt)
            }
         }
      }
   }
}

NPC9120013 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9120013(cm: cm))
   }
   return (NPC9120013) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }