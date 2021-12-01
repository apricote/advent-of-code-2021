package main

import "testing"

func TestGetMeasurementIncreases(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: `199
200
208
210
200
207
240
269
260
263`, want: 7},
		{input: `156
153
163
168
166
164
149
187
192
194
197
199
201
210
211
196
202
203
193
199
236
240
241
251
228
226
227
256
269
270
265
271
267
271
247
254
259
258
269
270
269
267
289
290
288
309
307
299
309
326
329
332
335
338
342
333
330
332
333
334
318
317
318
316
314
315
322
327
328
331
332
338
346
341
350
334
323
322
321
325
329
342
343
353
361
355
331
329
330
319
321
322
321
322
325
316
317
322
324
331
322
323
324
325
324
336
324
335
336
337
335
336
339
343
348
353
351
349
339
341
327
322
325
298
299
306
307
321
317
314
313
316
328
325
330
331
332
339
340
348
352
354
350
374
364
370
371
379
381
380
381
383
389
418
415
422
429
436
432
430
431
445
461
465
468
476
473
478
494
485
490
491
492
489
490
484
487
475
496
495
481
479
482
479
483
481
478
479
493
496
497
501
500
506
504
470
474
482
489
490
485
486
487
485
486
483
511
509
526
549
538
534
542
528
532
538
564
566
567
568
569
578
576
578
575
582
580
583
584
583
586
589
576
569
568
565
567
571
553
555
561
563
565
566
569
573
576
608
611
631
635
634
638
637
638
637
640
642
641
637
656
655
656
655
656
655
663
661
667
663
664
673
664
672
674
673
680
690
691
692
691
690
684
683
682
664
666
646
648
651
640
646
634
635
637
652
653
654
659
657
658
659
669
677
678
679
678
680
683
680
683
689
686
687
686
704
707
708
711
713
696
695
698
693
695
693
699
721
740
750
751
756
757
776
780
779
759
756
780
789
810
792
790
771
773
771
772
777
780
790
801
807
813
814
828
830
831
832
843
835
833
829
836
844
842
844
846
845
847
850
849
838
833
830
833
834
832
842
844
848
847
858
850
856
866
870
890
903
902
901
906
908
914
917
919
922
919
925
924
925
927
926
937
939
943
948
953
954
942
941
930
936
916
917
918
919
922
923
925
923
925
928
914
917
918
919
934
933
936
977
994
992
991
994
1017
1018
1015
1025
1007
987
975
973
979
978
1000
999
1000
1003
1009
1010
997
998
997
998
1002
1001
1007
1013
1020
1019
1031
1030
1032
1036
1041
1056
1044
1049
1062
1057
1056
1070
1073
1067
1066
1081
1059
1061
1053
1048
1056
1064
1067
1066
1064
1056
1055
1048
1054
1055
1042
1038
1039
1035
1050
1056
1057
1059
1069
1076
1070
1069
1075
1077
1071
1048
1037
1036
1051
1043
1049
1076
1075
1081
1080
1076
1075
1060
1067
1072
1099
1100
1104
1117
1119
1120
1098
1093
1092
1098
1097
1107
1112
1116
1122
1116
1120
1119
1107
1113
1127
1137
1145
1142
1144
1142
1143
1146
1143
1172
1187
1188
1163
1176
1188
1186
1201
1200
1208
1212
1215
1216
1226
1227
1232
1233
1245
1244
1245
1228
1227
1225
1226
1230
1240
1257
1260
1246
1247
1241
1240
1244
1251
1252
1254
1264
1266
1261
1262
1265
1276
1283
1293
1301
1295
1298
1302
1307
1308
1317
1318
1321
1336
1337
1327
1335
1357
1358
1348
1349
1364
1361
1364
1370
1362
1364
1367
1368
1360
1366
1368
1369
1371
1367
1329
1332
1334
1330
1325
1334
1341
1349
1354
1345
1344
1361
1362
1346
1369
1385
1383
1382
1383
1385
1382
1386
1406
1404
1393
1396
1409
1415
1419
1424
1423
1408
1405
1401
1381
1382
1376
1378
1383
1407
1412
1408
1405
1400
1398
1427
1446
1447
1432
1428
1454
1472
1476
1468
1482
1473
1466
1457
1446
1445
1471
1469
1468
1467
1465
1464
1471
1484
1478
1479
1492
1500
1487
1488
1499
1500
1512
1507
1521
1523
1533
1530
1521
1525
1530
1531
1533
1538
1542
1545
1549
1550
1553
1552
1551
1557
1564
1576
1577
1576
1573
1569
1568
1573
1561
1584
1588
1609
1606
1607
1614
1620
1634
1650
1651
1637
1636
1638
1643
1648
1649
1653
1652
1660
1656
1664
1669
1673
1692
1691
1679
1692
1695
1697
1700
1692
1690
1711
1738
1739
1743
1756
1752
1753
1751
1753
1748
1751
1749
1743
1742
1744
1760
1761
1759
1764
1765
1769
1778
1787
1788
1789
1788
1789
1791
1790
1789
1782
1783
1784
1776
1784
1785
1787
1810
1809
1821
1817
1818
1804
1803
1812
1792
1793
1798
1808
1812
1831
1834
1842
1817
1806
1825
1824
1820
1818
1823
1824
1826
1824
1821
1824
1809
1802
1790
1811
1814
1813
1814
1816
1813
1845
1857
1862
1861
1860
1861
1872
1892
1893
1877
1870
1869
1876
1865
1866
1840
1845
1853
1850
1848
1841
1840
1856
1877
1878
1874
1876
1871
1867
1866
1884
1893
1898
1899
1900
1898
1921
1926
1933
1940
1947
1937
1946
1947
1951
1964
1975
1985
1986
1974
1977
1971
1977
1979
1995
1983
1989
1991
2007
2006
2019
1999
1988
1994
1970
1981
1980
1967
1966
1965
1962
1956
1959
1957
1945
1943
1947
1945
1951
1952
1951
1952
1953
1938
1969
1965
1964
1975
1997
1998
2003
2005
2002
2012
2009
2017
2010
1994
2001
2000
1998
1997
1996
1993
1987
1986
2002
2003
2014
2022
2020
2021
2020
2027
2025
2026
2027
2026
2023
2020
2022
2032
2029
2017
2016
2005
2000
1993
1997
1996
1998
2016
2027
2025
2022
2027
2039
2045
2034
2037
2040
2019
2022
2024
2031
2030
2047
2044
2042
2061
2060
2061
2064
2063
2064
2065
2072
2070
2071
2074
2072
2054
2056
2055
2075
2078
2085
2081
2085
2086
2091
2072
2066
2067
2042
2053
2054
2088
2087
2101
2102
2106
2108
2106
2118
2116
2114
2120
2122
2101
2102
2113
2118
2119
2129
2132
2134
2130
2129
2131
2136
2124
2126
2144
2143
2141
2135
2132
2133
2130
2131
2127
2141
2165
2174
2173
2181
2163
2154
2160
2184
2180
2175
2179
2181
2180
2172
2169
2159
2153
2166
2169
2170
2180
2210
2218
2212
2234
2235
2236
2251
2243
2250
2251
2233
2213
2217
2219
2202
2203
2201
2199
2201
2203
2204
2205
2216
2195
2201
2200
2215
2223
2234
2236
2241
2250
2248
2247
2238
2239
2233
2232
2253
2248
2259
2261
2270
2259
2246
2245
2244
2241
2242
2247
2253
2258
2259
2261
2283
2285
2284
2292
2291
2292
2308
2301
2296
2303
2304
2303
2302
2301
2302
2324
2342
2343
2363
2388
2384
2385
2401
2416
2418
2414
2415
2393
2399
2382
2360
2361
2337
2345
2359
2373
2370
2371
2383
2384
2383
2387
2392
2396
2413
2419
2416
2424
2423
2434
2443
2469
2450
2438
2459
2461
2452
2461
2468
2479
2480
2482
2486
2485
2481
2485
2484
2486
2492
2501
2500
2501
2491
2494
2495
2498
2503
2504
2517
2511
2524
2528
2543
2546
2547
2540
2537
2540
2539
2546
2557
2561
2562
2546
2537
2549
2543
2553
2557
2562
2560
2559
2558
2574
2582
2577
2581
2569
2588
2589
2600
2611
2596
2597
2598
2595
2580
2576
2590
2591
2586
2603
2602
2608
2606
2637
2645
2653
2655
2651
2644
2676
2683
2688
2697
2700
2701
2688
2703
2715
2726
2727
2714
2725
2681
2683
2690
2685
2692
2693
2692
2693
2692
2697
2696
2690
2722
2721
2722
2721
2733
2735
2740
2733
2729
2733
2732
2733
2735
2747
2746
2747
2753
2749
2752
2758
2762
2772
2773
2775
2776
2783
2782
2784
2790
2784
2798
2797
2798
2788
2790
2789
2812
2813
2795
2796
2786
2792
2799
2768
2765
2759
2751
2753
2730
2729
2721
2720
2718
2722
2717
2701
2689
2682
2694
2682
2681
2682
2683
2662
2671
2670
2675
2674
2678
2685
2690
2681
2679
2685
2684
2670
2671
2688
2719
2721
2724
2723
2722
2753
2769
2771
2767
2765
2764
2763
2765
2769
2759
2760
2773
2790
2785
2788
2787
2781
2780
2795
2794
2795
2792
2823
2829
2837
2845
2852
2841
2844
2860
2855
2871
2888
2881
2876
2882
2886
2888
2906
2917
2925
2924
2926
2918
2922
2921
2922
2918
2905
2906
2916
2922
2902
2910
2892
2857
2853
2855
2852
2840
2850
2835
2837
2852
2860
2855
2860
2859
2868
2869
2883
2865
2864
2865
2866
2867
2874
2880
2886
2888
2898
2902
2892
2893
2886
2856
2853
2854
2851
2856
2857
2853
2862
2861
2862
2861
2866
2865
2851
2856
2836
2841
2846
2849
2843
2852
2842
2843
2844
2871
2874
2883
2867
2849
2858
2861
2868
2880
2905
2900
2907
2906
2904
2909
2923
2924
2923
2924
2926
2929
2921
2916
2919
2901
2899
2910
2912
2913
2915
2919
2922
2933
2932
2970
2996
2992
2988
2969
2950
2934
2935
2902
2905
2901
2904
2901
2893
2892
2893
2871
2872
2870
2887
2886
2888
2905
2910
2916
2918
2917
2919
2925
2906
2908
2934
2917
2930
2929
2930
2914
2909
2918
2919
2913
2904
2905
2902
2886
2894
2893
2896
2897
2900
2898
2899
2879
2876
2873
2876
2877
2881
2882
2883
2872
2869
2866
2867
2874
2867
2872
2873
2875
2874
2885
2883
2882
2870
2867
2868
2861
2865
2871
2893
2896
2932
2925
2932
2939
2938
2943
2938
2931
2944
2957
2958
2962
2943
2925
2914
2926
2932
2933
2936
2939
2940
2954
2952
2951
2905
2911
2907
2908
2914
2915
2913
2920
2928
2941
2949
2960
2936
2938
2937
2942
2945
2941
2959
2968
2943
2944
2931
2938
2939
2931
2914
2911
2915
2904
2917
2930
2931
2933
2926
2902
2914
2945
2944
2926
2933
2953
2955
2968
2966
2965
2967
2979
2957
2959
2957
2960
2944
2945
2949
2965
3013
3014
3021
3022
3037
3029
3024
3025
3024
3017
3013
3014
2996
2997
3011
3012
3014
3007
3032
3063
3058
3084
3083
3079
3069
3076
3077
3080
3073
3070
3071
3070
3060
3057
3058
3051
3050
3051
3040
3055
3060
3058
3059
3058
3057
3058
3057
3058
3057
3055
3065
3064
3084
3087
3085
3084
3082
3102
3101
3066
3067
3070
3085
3086
3091
3092
3084
3087
3083
3087
3085
3086
3089
3090
3074
3092
3089
3082
3076
3077
3071
3104
3101
3113
3118
3140
3148
3151
3160
3162
3171
3168
3173
3172
3205
3204
3206
3229
3236
3244
3245
3244
3280
3250
3251
3250
3249
3255
3256
3270
3271
3272
3276
3280
3278
3266
3272
3270
3272
3274
3293
3304
3309
3308
3305
3285
3280
3284
3293
3294
3296
3286
3300
3301
3295
3319
3324
3320
3321
3322
3323
3310
3294
3300
3299
3305
3322
3327
3339
3335
3336
3357
3379
3367
3359
3365
3386
3407
3403
3404
3402
3416
3391
3394
3381
3382
3385
3382
3375
3378
3379
3381
3380
3388
3404
3408
3404
3413
3421
3422
3423
3434
3440
3439
3444
3448
3458
3462
3459
3466
3465
3478
3477
3501
3498
3501
3500
3499
3513
3511
3490
3493
3495
3498
3496
3498
3499
3491
3502
3522
3523
3519
3522
3519
3514
3502
3498
3500
3484
3468
3480
3465
3468
3462
3465
3473
3474
3473
3480
3486
3482
3495
3470
3478
3457
3434
3469
3472
3503
3519
3526
3542
3541
3542
3552
3553
3568
3569
3564
3563
3565
3593
3591
3599
3612
3611
3613
3621
3620
3619
3622
3623
3633
3629
3630
3633
3634
3625
3618
3616
3617
3607
3632
3633
3637
3636
3643
3645
3666
3664
3677
3682
3686
3687
3694
3698
3715
3704
3709
3718
3716
3726
3727
3726
3725
3721
3705
3699
3702
3703
3696
3703
3704
3706
3710
3711
3701
3702
3692
3690
3669
3670
3682
3681
3688
3690
3676
3681
3684
3683
3697
3684
3694
3693
3692
3698
3723
3719
3745
3754
3755
3762
3769
3770`, want: 1233},
	}

	for _, tc := range tests {
		got := GetMeasurementIncreases(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d increases but got %d", tc.want, got)
		}
	}
}
