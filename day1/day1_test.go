/*
 * Copyright (c) 2020. Uriel Márquez All Rights Reserved
 * https://umarquez.c0d3.mx
 */

package day1_test

import (
	"log"
	"main/day1"
	"strconv"
	"strings"
	"testing"
)

var input = `1348
1621
1500
1818
1266
1449
1880
1416
1862
1665
1588
1704
1922
1482
1679
1263
1137
1045
1405
1048
1619
1520
455
1142
1415
1554
1690
1886
1891
1701
1915
1521
1253
1580
1376
1564
1747
1814
1749
1485
1969
974
1566
1413
1451
1200
1558
1756
1910
1044
470
1620
1772
1066
1261
1776
988
1976
1834
1896
1646
1626
1300
1692
1204
2006
1265
1911
1361
1766
1750
2000
1824
1726
1672
651
1226
1954
1055
1999
1793
1640
1567
1040
1426
1717
1658
1864
1917
695
1071
1573
1897
1546
1727
1801
1259
1290
1481
1148
1332
1262
1536
1184
1821
1681
1671
1612
1678
1703
1604
1697
2003
1453
1493
1797
1180
1234
1775
1859
1388
1393
667
1767
1429
1990
1322
1684
1696
1565
1380
1745
1685
1189
1396
1593
1850
1722
1495
1844
1285
1483
1635
1072
1947
1109
1586
1730
1723
1246
1389
1135
1827
1531
1583
1743
1958
183
1323
1949
1799
1269
1379
1950
1592
1467
1052
1418
2009
1227
1254
1865
1609
1848
1653
1691
1633
1349
1104
1790
1755
1847
1598
1872
1478
1778
1952
1694
1238
1825
1508
1141
1464
1838
1292
1403
1365
1494
934
1235`

var values []int

func TestFindNumbersAndMultiply(t *testing.T) {
	type args struct {
		sum    int
		nums   int
		values []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "incorrect numbers",
			args: args{
				sum:    2020,
				nums:   1,
				values: values,
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "two numbers that sums 2020",
			args: args{
				sum:    2020,
				nums:   2,
				values: values,
			},
			want:    712075,
			wantErr: false,
		},
		{
			name: "three numbers that sums 2020",
			args: args{
				sum:    2020,
				nums:   3,
				values: values,
			},
			want:    145245270,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day1.FindNumbersAndMultiply(tt.args.sum, tt.args.nums, tt.args.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindNumbersAndMultiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindNumbersAndMultiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func init() {
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("error decoding input text, %v", err)
		}
		values = append(values, num)
	}
}
