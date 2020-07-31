package repos

const (
	//插入platform SQL
	PLATFORM_INSERT_SQL = "INSERT INTO t_platform ( platform_code, platform_name, sign_key, gmt_create, gmt_modified, sign_type)  VALUES(?,?,?,now(),now(),?)"

	//查询platform SQL
	PLATFORM_QUERY_SQL = ""
)
