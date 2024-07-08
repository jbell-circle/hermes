# Hermes

A toy address book.
Designed for a workshop on PR stacking.

## Requirements

- golang v1.20

## Development

See Makefile for available commands.

## Stacking Pull Requests (PR) Tutorial

### Motivation

Smaller PRs are easier to review.
But you don't want to stop writing code while you wait for your PRs to be merged.

Stacking PRs allows you to keep working off a PR that's being reviewed.
It can be your own PR that you want to split into smaller, more manageable PRs.
Or it can be someone else's PR that you depend upon for your work.

### Getting Started

1. Fork this repo.
2. `git clone git@github.com:<$USERNAME>/hermes.git`

### Your task

Hermes is a microservice that is meant to store and retrieve addresses based on keys.

```sh
# create an address
curl -v 127.0.0.1:8080/addresses/:id -d '{"Name": "Justin Bell", "Address1": "99 High St", "City": "Boston", "State": "MA", "ZipCode": "02110"}'
```

This endpoint is already implemented.
Storing the data is implemented in [PR#1](https://github.com/jbell-circle/hermes/pull/1).

```sh
# retrieve an address
curl 127.0.0.1:8080/addresses/:id
```

Your job is to implement the GET endpoint.

### Working from existing branch

[PR#1](https://github.com/jbell-circle/hermes/pull/1) already has some of the changes we need to implement our feature but hasn't been merged yet.
You don't have to wait for these changes to be merged!

Add the author's fork as a remote and fetch all their branches.

```sh
git remote add jbell-circle git@github.com:jbell-circle/hermes.git
git fetch jbell-circle
```

Then you can switch to their branch and create your own feature branch from theirs.

```sh
git switch -c implementAddressModel jbell-circle/implementAddressModel
git checkout -b implementGetAddress
```

Now implement your feature!

test!!

### Q&A

- Q: What if the branch my changes are based on gets updated?
  - A: checkout the original branch, `git checkout <original-branch>`, and `git pull` the latest changes.
  - Then checkout back to your branch, `git checkout <your-branch>` and `git rebase <original-rebase>` the new changes on top of yours.

- Q: What if the branch my changes are based on get merged main/master?
  - A: we usually use "squash and merge" at Circle so there will be a new commit history. So you'll have to use the `--onto` flag to find the common ancestor. Suppose you have a `main` branch, a feature branch `f1`, and a feature branch `f2` that's based off `f1`. When `f1` gets squash merged into `main`, from `f2` you can then run `git rebease --onto main f1` to rebase onto the common ancestor.
