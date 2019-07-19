package headers

type CacheControlHeader struct {
     Max_age 				int 	`field:"max-age"`
     Max_stale 			int 	`field:"max-stale"`
     Min_fresh 			int 	`field:"min-fresh"`
     No_cache 				bool `field:"no-cache"`
     No_store 				bool `field:"no-store"`
     No_transform 			bool `field:"no-transform"`
     Only_if_cached 		bool `field:"only-if-cached"`
     Public 				bool `field:"public"`
     Private 				bool `field:"private"`
     Proxy_revalidate		bool `field:"proxy-revalidate"`
     S_max_age 			int 	`field:"s-max-age"`
     Immutable 			bool `field:"immutable"`
     Stale_while_revalidate 	int 	`field:"stale-while-revalidate"`
     Stale_if_error 		int 	`field:"stale-if-error"`
}
