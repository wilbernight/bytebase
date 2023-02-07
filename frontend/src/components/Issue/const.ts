import { CompatibilityErrorCode, SQLReviewPolicyErrorCode } from "@/types";

export const LocalizedSQLRuleErrorCodes = new Set<number>([
  SQLReviewPolicyErrorCode.STATEMENT_SYNTAX_ERROR,
  SQLReviewPolicyErrorCode.STATEMENT_NO_WHERE,
  SQLReviewPolicyErrorCode.STATEMENT_NO_SELECT_ALL,
  SQLReviewPolicyErrorCode.STATEMENT_LEADING_WILDCARD_LIKE,
  SQLReviewPolicyErrorCode.STATEMENT_CREATE_TABLE_AS,
  SQLReviewPolicyErrorCode.STATEMENT_DISALLOW_COMMIT,
  SQLReviewPolicyErrorCode.STATEMENT_REDUNDANT_ALTER_TABLE,
  SQLReviewPolicyErrorCode.STATEMENT_DML_DRY_RUN_FAILED,
  SQLReviewPolicyErrorCode.STATEMENT_AFFECTED_ROW_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.STATEMENT_ADD_COLUMN_WITH_DEFAULT,
  SQLReviewPolicyErrorCode.STATEMENT_ADD_CHECK_WITH_VALIDATION,
  SQLReviewPolicyErrorCode.STATEMENT_ADD_NOT_NULL,
  SQLReviewPolicyErrorCode.TABLE_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.COLUMN_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.INDEX_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.UK_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.FK_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.PK_NAMING_MISMATCH,
  SQLReviewPolicyErrorCode.NAMING_AUTO_INCREMENT_COLUMN_CONVENTION_MISMATCH,
  SQLReviewPolicyErrorCode.NO_REQUIRED_COLUMN,
  SQLReviewPolicyErrorCode.COLUMN_CANBE_NULL,
  SQLReviewPolicyErrorCode.CHANGE_COLUMN_TYPE,
  SQLReviewPolicyErrorCode.NOT_NULL_COLUMN_WITH_NO_DEFAULT,
  SQLReviewPolicyErrorCode.COLUMN_NOT_EXISTS,
  SQLReviewPolicyErrorCode.USE_CHANGE_COLUMN_STATEMENT,
  SQLReviewPolicyErrorCode.CHANGE_COLUMN_ORDER,
  SQLReviewPolicyErrorCode.NO_COLUMN_COMMENT,
  SQLReviewPolicyErrorCode.COLUMN_COMMENT_TOO_LONG,
  SQLReviewPolicyErrorCode.AUTO_INCREMENT_COLUMN_NOT_INTEGER,
  SQLReviewPolicyErrorCode.DISABLED_COLUMN_TYPE,
  SQLReviewPolicyErrorCode.COLUMN_EXISTS,
  SQLReviewPolicyErrorCode.DROP_ALL_COLUMNS,
  SQLReviewPolicyErrorCode.SET_COLUMN_CHARSET,
  SQLReviewPolicyErrorCode.CHAR_LENGTH_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.AUTO_INCREMENT_COLUMN_INITIAL_VALUE_NOT_MATCH,
  SQLReviewPolicyErrorCode.AUTO_INCREMENT_COLUMN_SIGNED,
  SQLReviewPolicyErrorCode.DEFAULT_CURRENT_TIME_COLUMN_COUNT_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.ON_UPDATE_CURRENT_TIME_COLUMN_COUNT_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.NO_DEFAULT,
  SQLReviewPolicyErrorCode.NOT_INNODB_ENGINE,
  SQLReviewPolicyErrorCode.NO_PK_IN_TABLE,
  SQLReviewPolicyErrorCode.FK_IN_TABLE,
  SQLReviewPolicyErrorCode.TABLE_DROP_NAMING_CONVENTION,
  SQLReviewPolicyErrorCode.TABLE_NOT_EXISTS,
  SQLReviewPolicyErrorCode.NO_TABLE_COMMENT,
  SQLReviewPolicyErrorCode.TABLE_COMMENT_TOO_LONG,
  SQLReviewPolicyErrorCode.TABLE_EXISTS,
  SQLReviewPolicyErrorCode.CREATE_TABLE_PARTITION,
  SQLReviewPolicyErrorCode.DATABASE_NOT_EMPTY,
  SQLReviewPolicyErrorCode.NOT_CURRENT_DATABASE,
  SQLReviewPolicyErrorCode.DATABASE_IS_DELETED,
  SQLReviewPolicyErrorCode.NOT_USE_INDEX,
  SQLReviewPolicyErrorCode.INDEX_KEY_NUMBER_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.INDEX_PK_TYPE,
  SQLReviewPolicyErrorCode.INDEX_TYPE_NO_BLOB,
  SQLReviewPolicyErrorCode.INDEX_EXISTS,
  SQLReviewPolicyErrorCode.PRIMARY_KEY_EXISTS,
  SQLReviewPolicyErrorCode.INDEX_EMPTY_KEYS,
  SQLReviewPolicyErrorCode.PRIMARY_KEY_NOT_EXISTS,
  SQLReviewPolicyErrorCode.INDEX_NOT_EXISTS,
  SQLReviewPolicyErrorCode.INCORRECT_INDEX_NAME,
  SQLReviewPolicyErrorCode.SPATIAL_INDEX_KEY_NULLABLE,
  SQLReviewPolicyErrorCode.DUPLICATE_COLUMN_IN_INDEX,
  SQLReviewPolicyErrorCode.INDEX_COUNT_EXCEEDS_LIMIT,
  SQLReviewPolicyErrorCode.CREATE_INDEX_UNCONCURRENTLY,
  SQLReviewPolicyErrorCode.DISABLED_CHARSET,
  SQLReviewPolicyErrorCode.INSERT_TOO_MANY_ROWS,
  SQLReviewPolicyErrorCode.UPDATE_USE_LIMIT,
  SQLReviewPolicyErrorCode.INSERT_USE_LIMIT,
  SQLReviewPolicyErrorCode.UPDATE_USE_ORDER_BY,
  SQLReviewPolicyErrorCode.DELETE_USE_ORDER_BY,
  SQLReviewPolicyErrorCode.DELETE_USE_LIMIT,
  SQLReviewPolicyErrorCode.INSERT_NOT_SPECIFY_COLUMN,
  SQLReviewPolicyErrorCode.INSERT_USE_ORDER_BY_RAND,
  SQLReviewPolicyErrorCode.DISABLED_COLLATION,
  SQLReviewPolicyErrorCode.COMMENT_TOO_LONG,
  CompatibilityErrorCode.COMPATIBILITY_DROP_DATABASE,
  CompatibilityErrorCode.COMPATIBILITY_RENAME_TABLE,
  CompatibilityErrorCode.COMPATIBILITY_DROP_TABLE,
  CompatibilityErrorCode.COMPATIBILITY_RENAME_COLUMN,
  CompatibilityErrorCode.COMPATIBILITY_DROP_COLUMN,
  CompatibilityErrorCode.COMPATIBILITY_ADD_PRIMARY_KEY,
  CompatibilityErrorCode.COMPATIBILITY_ADD_UNIQUE_KEY,
  CompatibilityErrorCode.COMPATIBILITY_ADD_FOREIGN_KEY,
  CompatibilityErrorCode.COMPATIBILITY_ADD_CHECK,
  CompatibilityErrorCode.COMPATIBILITY_ALTER_CHECK,
  CompatibilityErrorCode.COMPATIBILITY_ALTER_COLUMN,
]);