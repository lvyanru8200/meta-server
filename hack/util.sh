function util:create_gopath_tree() {
  local repo_root=$1
  local go_path=$2

  local go_pkg_dir="${go_path}/src/github.com/DataTunerX/meta-server"
  go_pkg_dir=$(dirname "${go_pkg_dir}")

  mkdir -p "${go_pkg_dir}"

  if [[ ! -e "${go_pkg_dir}" || "$(readlink "${go_pkg_dir}")" != "${repo_root}" ]]; then
    ln -snf "${repo_root}" "${go_pkg_dir}"
  fi
}