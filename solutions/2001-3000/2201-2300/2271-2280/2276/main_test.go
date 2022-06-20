package mario

import "testing"

func TestConstructor(t *testing.T) {
	ins := Constructor()

	expected := []int{525, 525, 846, 846, 846, 847, 847, 847, 847, 924, 924, 955, 965, 965, 965, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 988, 988, 988, 988, 988, 988, 988, 1000, 1000, 1000, 1000, 1000}
	res := make([]int, 0, len(expected))

	{
		ins.Add(464, 988)
		ins.Add(720, 795)
		res = append(res, ins.Count())
		ins.Add(567, 771)
		res = append(res, ins.Count())
		ins.Add(143, 893)
		res = append(res, ins.Count())
		ins.Add(259, 410)
		ins.Add(461, 600)
		ins.Add(759, 771)
		ins.Add(944, 969)
		res = append(res, ins.Count())
		ins.Add(821, 908)
		ins.Add(695, 728)
		res = append(res, ins.Count())
		ins.Add(332, 373)
		ins.Add(369, 977)
		ins.Add(1000, 1000)
		res = append(res, ins.Count())
		ins.Add(319, 565)
		res = append(res, ins.Count())
		ins.Add(283, 786)
		ins.Add(191, 439)
		res = append(res, ins.Count())
		ins.Add(583, 851)
		res = append(res, ins.Count())
		ins.Add(374, 799)
		ins.Add(66, 264)
		ins.Add(793, 977)
		res = append(res, ins.Count())
		ins.Add(104, 689)
		res = append(res, ins.Count())
		ins.Add(301, 394)
		ins.Add(35, 632)
		ins.Add(500, 969)
		ins.Add(833, 907)
		ins.Add(113, 492)
		ins.Add(898, 954)
		res = append(res, ins.Count())
		ins.Add(724, 998)
		ins.Add(465, 764)
		res = append(res, ins.Count())
		ins.Add(368, 465)
		res = append(res, ins.Count())
		ins.Add(246, 391)
		ins.Add(99, 456)
		res = append(res, ins.Count())
		ins.Add(287, 377)
		ins.Add(996, 999)
		ins.Add(906, 937)
		res = append(res, ins.Count())
		ins.Add(972, 978)
		ins.Add(375, 579)
		res = append(res, ins.Count())
		ins.Add(491, 549)
		ins.Add(762, 766)
		ins.Add(158, 497)
		ins.Add(819, 829)
		res = append(res, ins.Count())
		ins.Add(899, 962)
		res = append(res, ins.Count())
		ins.Add(287, 761)
		res = append(res, ins.Count())
		ins.Add(349, 778)
		ins.Add(476, 574)
		res = append(res, ins.Count())
		ins.Add(992, 994)
		ins.Add(672, 903)
		ins.Add(462, 842)
		res = append(res, ins.Count())
		ins.Add(412, 1000)
		res = append(res, ins.Count())
		ins.Add(453, 768)
		res = append(res, ins.Count())
		ins.Add(774, 777)
		res = append(res, ins.Count())
		ins.Add(237, 436)
		res = append(res, ins.Count())
		ins.Add(973, 976)
		res = append(res, ins.Count())
		ins.Add(804, 975)
		res = append(res, ins.Count())
		ins.Add(77, 701)
		ins.Add(399, 420)
		res = append(res, ins.Count())
		ins.Add(227, 367)
		res = append(res, ins.Count())
		ins.Add(414, 665)
		ins.Add(936, 967)
		ins.Add(296, 654)
		ins.Add(270, 585)
		res = append(res, ins.Count())
		ins.Add(852, 860)
		ins.Add(456, 983)
		ins.Add(151, 970)
		res = append(res, ins.Count())
		ins.Add(13, 902)
		res = append(res, ins.Count())
		ins.Add(471, 565)
		res = append(res, ins.Count())
		ins.Add(403, 861)
		ins.Add(777, 921)
		res = append(res, ins.Count())
		ins.Add(968, 977)
		res = append(res, ins.Count())
		ins.Add(974, 984)
		res = append(res, ins.Count())
		ins.Add(637, 955)
		res = append(res, ins.Count())
		ins.Add(372, 678)
		res = append(res, ins.Count())
		ins.Add(1, 154)
		res = append(res, ins.Count())
		ins.Add(590, 670)
		ins.Add(10, 659)
		res = append(res, ins.Count())
		ins.Add(938, 950)
		res = append(res, ins.Count())
		ins.Add(200, 255)
		res = append(res, ins.Count())
		ins.Add(784, 965)
		ins.Add(723, 976)
		ins.Add(160, 688)
		ins.Add(265, 714)
		res = append(res, ins.Count())
		ins.Add(581, 887)
	}

	if len(expected) != len(res) {
		t.Logf("unmatched length, get: %d, want: %d\n", len(res), len(expected))
		t.Fail()
	}

	for i := range res {
		if res[i] != expected[i] {
			t.Logf("test failed on case 1, index: %d, get: %d, want: %d\n", i, res[i], expected[i])
			t.Fail()
		}
	}
}