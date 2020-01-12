
function yam --description "YAM credential manager"
    set -l cmd $argv[1]
    if test $cmd = "activate"
        set -l activate_args
        if test (count $argv) -gt 1
            set activate_args $argv[2..-1]
        end
        eval (__yam-helper __get $activate_args)
    else
        # send everything to helper
        __yam-helper $argv
    end
end
