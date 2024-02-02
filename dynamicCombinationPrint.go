package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	trigger := false
	if n > 0 {
		for firstLoop := 0; firstLoop <= 9; firstLoop++ {
			if n > 1 {
				for secondLoop := firstLoop + 1; secondLoop <= 9; secondLoop++ {
					if n > 2 {
						for thirdLoop := secondLoop + 1; thirdLoop <= 9; thirdLoop++ {
							if n > 3 {
								for fourthLoop := thirdLoop + 1; fourthLoop <= 9; fourthLoop++ {
									if n > 4 {
										for fifthLoop := fourthLoop + 1; fifthLoop <= 9; fifthLoop++ {
											if n > 5 {
												for sixthLoop := fifthLoop + 1; sixthLoop <= 9; sixthLoop++ {
													if n > 6 {
														for seventhLoop := sixthLoop + 1; seventhLoop <= 9; seventhLoop++ {
															if n > 7 {
																for eightLoop := seventhLoop + 1; eightLoop <= 9; eightLoop++ {
																	if n > 8 {
																		for ninthLoop := eightLoop + 1; ninthLoop <= 9; ninthLoop++ {
																			if trigger {
																				z01.PrintRune(',')
																				z01.PrintRune(' ')
																			}
																			z01.PrintRune(rune(firstLoop + '0'))
																			z01.PrintRune(rune(secondLoop + '0'))
																			z01.PrintRune(rune(thirdLoop + '0'))
																			z01.PrintRune(rune(fourthLoop + '0'))
																			z01.PrintRune(rune(fifthLoop + '0'))
																			z01.PrintRune(rune(sixthLoop + '0'))
																			z01.PrintRune(rune(seventhLoop + '0'))
																			z01.PrintRune(rune(eightLoop + '0'))
																			z01.PrintRune(rune(ninthLoop + '0'))
																			trigger = true
																		}
																	}
																	if n < 9 {
																		if trigger {
																			z01.PrintRune(',')
																			z01.PrintRune(' ')
																		}
																		z01.PrintRune(rune(firstLoop + '0'))
																		z01.PrintRune(rune(secondLoop + '0'))
																		z01.PrintRune(rune(thirdLoop + '0'))
																		z01.PrintRune(rune(fourthLoop + '0'))
																		z01.PrintRune(rune(fifthLoop + '0'))
																		z01.PrintRune(rune(sixthLoop + '0'))
																		z01.PrintRune(rune(seventhLoop + '0'))
																		z01.PrintRune(rune(eightLoop + '0'))
																		trigger = true
																	}
																}
															}
															if n < 8 {
																if trigger {
																	z01.PrintRune(',')
																	z01.PrintRune(' ')
																}
																z01.PrintRune(rune(firstLoop + '0'))
																z01.PrintRune(rune(secondLoop + '0'))
																z01.PrintRune(rune(thirdLoop + '0'))
																z01.PrintRune(rune(fourthLoop + '0'))
																z01.PrintRune(rune(fifthLoop + '0'))
																z01.PrintRune(rune(sixthLoop + '0'))
																z01.PrintRune(rune(seventhLoop + '0'))
																trigger = true
															}
														}
													}
													if n < 7 {
														if trigger {
															z01.PrintRune(',')
															z01.PrintRune(' ')
														}
														z01.PrintRune(rune(firstLoop + '0'))
														z01.PrintRune(rune(secondLoop + '0'))
														z01.PrintRune(rune(thirdLoop + '0'))
														z01.PrintRune(rune(fourthLoop + '0'))
														z01.PrintRune(rune(fifthLoop + '0'))
														z01.PrintRune(rune(sixthLoop + '0'))
														trigger = true
													}
												}
											}
											if n < 6 {
												if trigger {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}
												z01.PrintRune(rune(firstLoop + '0'))
												z01.PrintRune(rune(secondLoop + '0'))
												z01.PrintRune(rune(thirdLoop + '0'))
												z01.PrintRune(rune(fourthLoop + '0'))
												z01.PrintRune(rune(fifthLoop + '0'))
												trigger = true
											}
										}
									}
									if n < 5 {
										if trigger {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
										z01.PrintRune(rune(firstLoop + '0'))
										z01.PrintRune(rune(secondLoop + '0'))
										z01.PrintRune(rune(thirdLoop + '0'))
										z01.PrintRune(rune(fourthLoop + '0'))
										trigger = true
									}
								}
							}
							if n < 4 {
								if trigger {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
								z01.PrintRune(rune(firstLoop + '0'))
								z01.PrintRune(rune(secondLoop + '0'))
								z01.PrintRune(rune(thirdLoop + '0'))
								trigger = true
							}
						}
					}
					if n < 3 {
						if trigger {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune(firstLoop + '0'))
						z01.PrintRune(rune(secondLoop + '0'))
						trigger = true
					}
				}
			}
			if n < 2 {
				if trigger {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(firstLoop + '0'))
				trigger = true
			}
		}
	}
	z01.PrintRune('\n')
}
