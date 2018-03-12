#!/usr/bin/env bats

load test_helper

setup() {
  global_setup
  create_app
}

teardown() {
  destroy_app
  global_teardown
}

@test "(git) deploy specific branch" {
  run ssh dokku@dokku.me git:set --global deploy-branch global-branch
  GIT_REMOTE_BRANCH=global-branch deploy_app
  echo "output: "$output
  echo "status: "$status
  assert_success

  run ssh dokku@dokku.me ps:rebuild $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success

  run ssh dokku@dokku.me git:set $TEST_APP deploy-branch  app-branch
  GIT_REMOTE_BRANCH=app-branch deploy_app
  echo "output: "$output
  echo "status: "$status
  assert_success

  run ssh dokku@dokku.me ps:rebuild $TEST_APP
  echo "output: "$output
  echo "status: "$status
  assert_success

  run ssh dokku@dokku.me git:set --global deploy-branch
}
