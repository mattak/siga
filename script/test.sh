#!/bin/sh

setup() {
  rm -rf /tmp/test
  mkdir -p /tmp/test/result
}

teardown() {
  rm -rf /tmp/test
}

run_test() {
  input_file=$1
  expect_file=$2
  shift
  shift

  command_format="$* < $input_file"
  "$@" < $input_file > /tmp/test/result.tsv
  diff /tmp/test/result.tsv $expect_file > /tmp/test/diff.patch
  if [ $? -ne 0 ]; then
    cat << __FAILED__
❌ $command_format
- input: $input_file
- expect: $expect_file

diff <($command_format) $expect_file
__FAILED__
    cat /tmp/test/diff.patch
    exit 1
  else
    echo "✅ $command_format"
  fi
}

setup

INPUT_DIR=./test/resource/input
EXPECT_DIR=./test/resource/expect

# lt
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga lt -l expect value 0
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga lt -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga lt -l expect value 2

# le
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga le -l expect value 0
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga le -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga le -l expect value 2

# ge
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga ge -l expect value 0
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga ge -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga ge -l expect value 2

# gt
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga gt -l expect value 0
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga gt -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga gt -l expect value 2

# eq
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga eq -l expect value 0
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga eq -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga eq -l expect value 2

# filter or
run_test $INPUT_DIR/filter_123.tsv $EXPECT_DIR/filter_or_123.tsv siga filter or 1 value1 value2
run_test $INPUT_DIR/filter_123.tsv $EXPECT_DIR/filter_nor_123.tsv siga filter nor 1 value1 value2
run_test $INPUT_DIR/filter_123.tsv $EXPECT_DIR/filter_and_123.tsv siga filter and 1 value1 value2
run_test $INPUT_DIR/filter_123.tsv $EXPECT_DIR/filter_nand_123.tsv siga filter nand 1 value1 value2

# reverse
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/reverse_123.tsv siga reverse

# take
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/take_1.tsv siga take 1

# reverse_take
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/retake_3.tsv siga retake 1

# shift
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/shift_1.tsv siga shift -l expect value 1

# detect_any
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/expect_bare_true.tsv siga detect_any value 1
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/expect_bare_false.tsv siga detect_any value 0

# detect_all
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/expect_bare_false.tsv siga detect_all value 1
run_test $INPUT_DIR/multi_111.tsv $EXPECT_DIR/expect_bare_true.tsv siga detect_all value 1

# count
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/count_1.tsv siga count

# select_last_value
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/expect_bare_3.tsv siga select_last_value value

# sma
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/sma_123.tsv siga sma value 3

# hma
run_test $INPUT_DIR/multi_122.tsv $EXPECT_DIR/hma_122.tsv siga hma value 3

# identity
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/id_1.tsv siga id

# rename
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/rename_1.tsv siga rc value new_value

# multiply
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga mul -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga mul -l expect value 0
run_test $INPUT_DIR/single_0.tsv $EXPECT_DIR/expect_0_single_0.tsv siga mul -l expect value 1

# divide
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_1_single_1.tsv siga div -l expect value 1
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_nan_single_1.tsv siga div -l expect value 0
run_test $INPUT_DIR/single_0.tsv $EXPECT_DIR/expect_0_single_0.tsv siga div -l expect value 1
run_test $INPUT_DIR/single_nan.tsv $EXPECT_DIR/expect_nan_single_nan.tsv siga div -l expect value 1

# add
run_test $INPUT_DIR/single_0.tsv $EXPECT_DIR/expect_1_single_0.tsv siga add -l expect value 1

# sub
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/expect_0_single_1.tsv siga sub -l expect value 1

# sum
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/sum_123.tsv siga sum value 3

# const
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/const_123.tsv siga const -l expect 1

# deviations
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/deviations_123_2.tsv siga dev -l expect value 2
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/deviations_123_3.tsv siga dev -l expect value 3

# profit factor
run_test $INPUT_DIR/profit_factor_123.tsv $EXPECT_DIR/expect_bare_1.tsv siga pf value

# payoff ratio
run_test $INPUT_DIR/payoff_ratio_111.tsv $EXPECT_DIR/payoff_ratio_111.tsv siga po value

# trading evaluation
run_test $INPUT_DIR/buy_sell_1.tsv $EXPECT_DIR/buy_sell_1.tsv siga te buy sell

# dollar cost average
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/dollar_cost_average_123.tsv siga dca value

# normalize
run_test $INPUT_DIR/multi_421.tsv $EXPECT_DIR/normalize_421.tsv siga normal value

# gen const
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/gen_const_111.tsv siga gen const -l expect 3 1

# gen rand
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/gen_const_111.tsv siga gen rand -l expect 3 1 1

# gen rwalk
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/gen_const_111.tsv siga gen rwalk -l expect 3 1 1:0:0

teardown
