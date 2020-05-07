name = "fs"

service {
}

storage {
  op "list_dir" {
    optional = ["dir_func","file_func"]
  }
  op "new" {
    optional = ["work_dir"]
  }
  op "read" {
    optional = ["offset", "size"]
  }
  op "write" {
    required = null
    optional = ["size"]
  }
}