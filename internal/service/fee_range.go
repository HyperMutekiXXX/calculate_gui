package service

import (
	"calculate_gui/constant"
	"calculate_gui/internal/model"
)

var FeeRangeType = map[int32][]model.FeeRange{
	constant.GoodsBidding: {
		{0, 0.015},
		{1000000, 0.011},
		{5000000, 0.008},
		{10000000, 0.005},
		{50000000, 0.0025},
		{100000000, 0.0005},
		{1000000000, 0.0001},
	},
	constant.ServiceBidding: {
		{0, 0.015},
		{1000000, 0.008},
		{5000000, 0.0045},
		{10000000, 0.0025},
		{50000000, 0.001},
		{100000000, 0.0005},
		{1000000000, 0.0001},
	},
	constant.EngineeringBidding: {
		{0, 0.01},
		{1000000, 0.007},
		{5000000, 0.0055},
		{10000000, 0.0035},
		{50000000, 0.002},
		{100000000, 0.0005},
		{1000000000, 0.0001},
	},
}

func calculateBidAgentFee(amount float64, feeRanges []model.FeeRange) float64 {
	totalFee := 0.0
	len := len(feeRanges)

	for i := 0; i < len-1; i++ {
		if amount > feeRanges[i+1].UpperLimit {
			totalFee += (feeRanges[i+1].UpperLimit - feeRanges[i].UpperLimit) * feeRanges[i].Rate
		} else {
			totalFee += (amount - feeRanges[i].UpperLimit) * feeRanges[i].Rate
			return totalFee
		}
	}

	totalFee += (amount - feeRanges[len-1].UpperLimit) * feeRanges[len-1].Rate

	return totalFee
}

func CalFee(amount float64, feeType int32) float64 {
	return calculateBidAgentFee(amount, FeeRangeType[feeType])
}
