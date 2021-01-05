package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201123 {
    NPCConversationManager cm
    int status = -1
    int sel = -1

    int map = 102000003
    String job = "Warrior"
    int jobType = 1
    String no = "Come back to me if you decided to be a #b" + job + "#k."

    def start() {
        status = -1
        action((byte) 1, (byte) 0, 0)
    }

    def action(Byte mode, Byte type, Integer selection) {
        if (mode == -1) {
            cm.sendOk(no)
            cm.dispose()
        } else {
            if (mode == 0 && type > 0) {
                cm.sendOk(no)
                cm.dispose()
            }

            if (mode == 1) {
                status++
            } else {
                status--
            }

            if (status == 0) {
                if (cm.getJobId() == MapleJob.BEGINNER.getId()) {
                    if (cm.getLevel() >= 10 && cm.canGetFirstJob(jobType)) {
                        cm.sendYesNo("9201123_DO_YOU_WANT_TO_GO", map, job)
                    } else {
                        cm.sendOk("9201123_IF_YOU_WANT", job, cm.getFirstJobStatRequirement(jobType))
                        cm.dispose()
                    }
                } else {
                    cm.sendOk("9201123_MUCH_STRONGER_NOW")
                    cm.dispose()
                }
            } else if (status == 1) {
                cm.warp(map, 0)
                cm.dispose()
            }
        }
    }
}

NPC9201123 getNPC() {
    if (!getBinding().hasVariable("npc")) {
        NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
        getBinding().setVariable("npc", new NPC9201123(cm: cm))
    }
    return (NPC9201123) getBinding().getVariable("npc")
}

def start() {
    getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }