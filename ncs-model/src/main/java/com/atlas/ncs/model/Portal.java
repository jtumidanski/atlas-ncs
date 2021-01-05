package com.atlas.ncs.model;

public record Portal(int id, String name, String target, int type, int x, int y, int targetMap, String scriptName) {
}
