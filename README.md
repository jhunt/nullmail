nullmail - A Null Router for SMTP
=================================

`nullmail` This is a very small, standalone binary that listens on
a given port, speaks unencrypted SMTP, and dumps all the mail it
receives to ... nowhere.  You do get logs though!

Usage is simple:

```
$ nullmail
listening for inbound SMTP on :2525
>> received mail from from@foo for to@bar with subject: xyzzy
```

If you want a different port:

```
$ PORT=5587 nullmail
listening for inbound SMTP on :5587
```

If you want to save the logs:

```
$ nullmail | tee /var/log/nullmail.log
```

If you don't want the logs:

```
$ nullmail >/dev/null
```

etc.
