use JSON;
my %rec_hash = ('message' => 'Hello World!');
my $json = encode_json \%rec_hash;
print "$json";