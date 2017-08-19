CREATE VIEW "poolmate"."session_sets" AS
  SELECT
    t.id,
    t.deleted,
    t.version,
    t.session_id,
    t.set_number,
    t.swimming_time,
    t.rest_time,
    t.swimming_time + t.rest_time AS total_time,
    t.laps,
    t.laps * s.pool_length        AS distance,
    t.average_strokes,
    t.speed,
    t.efficiency_index
  FROM poolmate.session_set AS t
    INNER JOIN poolmate.session AS s ON t.session_id = s.id
  ORDER BY t.id;