// Code generated by "stringer -type=MeanOfDeath"; DO NOT EDIT.

package quake3

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MOD_UNKNOWN-0]
	_ = x[MOD_SHOTGUN-1]
	_ = x[MOD_GAUNTLET-2]
	_ = x[MOD_MACHINEGUN-3]
	_ = x[MOD_GRENADE-4]
	_ = x[MOD_GRENADE_SPLASH-5]
	_ = x[MOD_ROCKET-6]
	_ = x[MOD_ROCKET_SPLASH-7]
	_ = x[MOD_PLASMA-8]
	_ = x[MOD_PLASMA_SPLASH-9]
	_ = x[MOD_RAILGUN-10]
	_ = x[MOD_LIGHTNING-11]
	_ = x[MOD_BFG-12]
	_ = x[MOD_BFG_SPLASH-13]
	_ = x[MOD_WATER-14]
	_ = x[MOD_SLIME-15]
	_ = x[MOD_LAVA-16]
	_ = x[MOD_CRUSH-17]
	_ = x[MOD_TELEFRAG-18]
	_ = x[MOD_FALLING-19]
	_ = x[MOD_SUICIDE-20]
	_ = x[MOD_TARGET_LASER-21]
	_ = x[MOD_TRIGGER_HURT-22]
	_ = x[MOD_NAIL-23]
	_ = x[MOD_CHAINGUN-24]
	_ = x[MOD_PROXIMITY_MINE-25]
	_ = x[MOD_KAMIKAZE-26]
	_ = x[MOD_JUICED-27]
	_ = x[MOD_GRAPPLE-28]
}

const _MeanOfDeath_name = "MOD_UNKNOWNMOD_SHOTGUNMOD_GAUNTLETMOD_MACHINEGUNMOD_GRENADEMOD_GRENADE_SPLASHMOD_ROCKETMOD_ROCKET_SPLASHMOD_PLASMAMOD_PLASMA_SPLASHMOD_RAILGUNMOD_LIGHTNINGMOD_BFGMOD_BFG_SPLASHMOD_WATERMOD_SLIMEMOD_LAVAMOD_CRUSHMOD_TELEFRAGMOD_FALLINGMOD_SUICIDEMOD_TARGET_LASERMOD_TRIGGER_HURTMOD_NAILMOD_CHAINGUNMOD_PROXIMITY_MINEMOD_KAMIKAZEMOD_JUICEDMOD_GRAPPLE"

var _MeanOfDeath_index = [...]uint16{0, 11, 22, 34, 48, 59, 77, 87, 104, 114, 131, 142, 155, 162, 176, 185, 194, 202, 211, 223, 234, 245, 261, 277, 285, 297, 315, 327, 337, 348}

func (i MeanOfDeath) String() string {
	if i < 0 || i >= MeanOfDeath(len(_MeanOfDeath_index)-1) {
		return "MeanOfDeath(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MeanOfDeath_name[_MeanOfDeath_index[i]:_MeanOfDeath_index[i+1]]
}