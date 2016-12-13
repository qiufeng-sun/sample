class xbox {
  include config

  file { [  "/home/work/bin/",
            "${config::log_base}" ] :
    ensure   => directory,
    owner    => "${config::user}",
    group    => "${config::user}",
    mode     => 0755,
  } ->
  package { 'cronolog':
    ensure   => 'installed',
    provider => 'yum',
  }
 
}

include xbox
