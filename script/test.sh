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

# count
run_test $INPUT_DIR/single_1.tsv $EXPECT_DIR/count_1.tsv siga count

# sum
run_test $INPUT_DIR/multi_123.tsv $EXPECT_DIR/sum_123.tsv siga sum value 3

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

teardown
