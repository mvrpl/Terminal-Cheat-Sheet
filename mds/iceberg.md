#### PROCEDURES
|||
|-|-|
`CALL catalog_name.system.set_current_snapshot(table => 'db.sample', ref => 's1');`|Set the current snapshot for db.sample to tag s1
`CALL catalog_name.system.cherrypick_snapshot(snapshot_id => 1, table => 'my_table' );`|Cherry-pick snapshot 1 with named args
`CALL hive_prod.system.expire_snapshots('db.sample', TIMESTAMP '2021-06-30 00:00:00.000', 100);`|Remove snapshots older than specific day and time, but retain the last 100 snapshots
`CALL hive_prod.system.expire_snapshots(table => 'db.sample', snapshot_ids => ARRAY(123));`|Remove snapshots with snapshot ID 123 (note that this snapshot ID should not be the current snapshot)
`CALL catalog_name.system.remove_orphan_files(table => 'db.sample', dry_run => true);`|List all the files that are candidates for removal by performing a dry run of the remove_orphan_files command on this table without actually removing them
`CALL catalog_name.system.set_current_snapshot('db.sample', 1);`|Set the current snapshot for db.sample to 1
`CALL catalog_name.system.rollback_to_timestamp('db.sample', TIMESTAMP '2021-06-30 00:00:00.000');`|Roll back db.sample to a specific day and time.
`CALL catalog_name.system.cherrypick_snapshot('my_table', 1);`|Cherry-pick snapshot 1
`CALL catalog_name.system.remove_orphan_files(table => 'db.sample', location => 'tablelocation/data');`|Remove any files in the tablelocation/data folder which are not known to the table db.sample
`CALL catalog_name.system.rewrite_data_files(table => 'db.sample', where => 'id = 3 and name = "foo"');`|Rewrite the data files in table db.sample and select the files that may contain data matching the filter (id = 3 and name = "foo") to be rewritten
`CALL catalog_name.system.rollback_to_snapshot('db.sample', 1);`|Roll back table db.sample to snapshot ID 1
