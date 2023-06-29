package biometric_service

import (
	cv2 "gocv.io/x/gocv"
)

func FingerPrintMatch(targetPath string, samplePath string) float32 {
	target := cv2.IMRead(targetPath, cv2.IMReadGrayScale)
	sample := cv2.IMRead(samplePath, cv2.IMReadGrayScale)

	sift := cv2.NewSIFT()

	targetKeypoints, targetDescriptors := sift.DetectAndCompute(target, cv2.NewMat())
	_, sampleDescriptors := sift.DetectAndCompute(sample, cv2.NewMat())

	matcher := cv2.NewBFMatcher()
	matches := matcher.KnnMatch(targetDescriptors, sampleDescriptors, 2)

	var goodMatches []cv2.DMatch
	for _, match := range matches {
		if len(match) >= 2 {
			if match[0].Distance < 0.75*match[1].Distance {
				goodMatches = append(goodMatches, match[0])
			}
		}
	}

	return float32(len(goodMatches)) / float32(len(targetKeypoints)) * 100
}
